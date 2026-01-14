package core

type ApiFormat string

const (
	JSON ApiFormat = "json"
	XML  ApiFormat = "xml"
)

type SignMethod string

const (
	MD5  SignMethod = "md5"
	HMAC SignMethod = "hmac"
)
