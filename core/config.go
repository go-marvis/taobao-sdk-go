package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	AppKey     string     // 应用标识
	Secret     string     // 请求加密
	BaseUrl    string     // 服务地址
	SignMethod SignMethod // 签名摘要算法
	Format     ApiFormat  // 响应格式
	Version    string     // API协议版本
	Session    string     // 用户登录授权信息

	Debug bool // 调试模式, 日志输出请求和响应

	Retry   int
	Limiter *rate.Limiter

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
