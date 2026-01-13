package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	JSON ApiFormat = "json"
	XML  ApiFormat = "xml"
)

type SignMethod string

const (
	MD5  SignMethod = "md5"
	HMAC SignMethod = "hmac"
)

type Config struct {
	AppKey     string
	Secret     string
	BaseUrl    string
	SignMethod SignMethod
	Format     ApiFormat
	Version    string
	Session    string

	Retry   int
	Limiter *rate.Limit

	HttpClient HttpClient
	ReqTimeout time.Duration
	Serializer Serializer
}

func Init(config *Config) {
	if config.HttpClient == nil {
		if config.ReqTimeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.ReqTimeout}
		}
	}

	if config.Serializer == nil {
		config.Serializer = &defaultSerializer{}
	}
}
