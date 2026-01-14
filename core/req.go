package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod  string
	Method      string
	QueryParams url.Values
}

func NewApiReq() *ApiReq {
	return &ApiReq{
		QueryParams: url.Values{},
	}
}

type ReqOption struct {
	Header http.Header

	TargetAppKey string
	Session      string
	PartnerId    string
}

type ReqOptionFunc func(option *ReqOption)

func WithSession(session string) ReqOptionFunc {
	return func(option *ReqOption) {
		option.Session = session
	}
}
