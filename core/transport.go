package core

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"
)

const (
	defaultContentType = "application/x-www-form-urlencoded;charset=utf-8"
)

func doSend(client HttpClient, req *http.Request) (*ApiResp, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ApiResp{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}

func readResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Request(ctx context.Context, apiReq *ApiReq, config *Config, options ...ReqOptionFunc) (*ApiResp, error) {
	values := make(url.Values, 0)
	for key := range apiReq.QueryParams {
		values.Set(key, apiReq.QueryParams.Get(key))
	}

	values.Set("method", apiReq.Method)
	values.Set("app_key", config.AppKey)
	values.Set("sign_method", string(config.SignMethod))
	values.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	values.Set("format", string(config.Format))
	values.Set("v", config.Version)
	values.Set("session", config.Session)

	option := &ReqOption{}
	for _, optf := range options {
		optf(option)
	}
	if option.TargetAppKey != "" {
		values.Set("target_app_key", option.TargetAppKey)
	}
	if option.Session != "" {
		values.Set("session", option.Session)
	}
	if option.PartnerId != "" {
		values.Set("partner_id", option.PartnerId)
	}
	if config.Format == "json" {
		values.Set("simplify", "true")
	}

	values.Set("sign", sign(values, config))

	url := config.BaseUrl
	payload := values.Encode()
	if config.Debug {
		log.Printf("[DEBUG] [API] %s %s payload:\n%s\n", apiReq.HttpMethod, url, payload)
	}

	req, err := http.NewRequestWithContext(ctx, apiReq.HttpMethod, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", defaultContentType)
	for k, vs := range option.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	var apiResp *ApiResp
	for range config.Retry {
		if config.Limiter != nil {
			if err = config.Limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		apiResp, err = doSend(config.HttpClient, req)
		if apiResp.StatusCode != http.StatusOK {
			slog.Error("http error", "status", apiResp.StatusCode)
			return apiResp, fmt.Errorf("http error, status=%d", apiResp.StatusCode)
		}

		errResp := &ErrorResponse{}
		if err := config.Serializer.Unmarshal(apiResp.RawBody, errResp); err != nil {
			slog.Error("unmarshal error", "err", err)
			return apiResp, err
		}
		if errResp.Code != 0 {
			slog.Error("api error", "code", errResp.Code, "msg", errResp.Msg)
			time.Sleep(2 * time.Second)
			continue
		}

		err = nil
		break
	}
	return apiResp, err
}

func sign(params url.Values, config *Config) string {
	var buf bytes.Buffer
	if config.SignMethod == MD5 {
		buf.WriteString(config.Secret)
	}
	for _, key := range slices.Sorted(maps.Keys(params)) {
		buf.WriteString(key)
		buf.WriteString(params.Get(key))
	}
	if config.SignMethod == MD5 {
		buf.WriteString(config.Secret)
	}

	var h hash.Hash
	switch config.SignMethod {
	case MD5:
		h = md5.New()
	case HMAC:
		h = hmac.New(md5.New, []byte(config.Secret))
	default:
		panic("missing sign_method")
	}
	io.Copy(h, bytes.NewReader(buf.Bytes()))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
