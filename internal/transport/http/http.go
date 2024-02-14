package http_client

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var MaxConnPerHost = &env.Env.MaxConnsPerHost
var MaxConnWaitTimeout = &env.Env.MaxConnWaitTimeout
var ReadTimeout = &env.Env.ReadTimeout
var MaxIdempotentCallAttempts = &env.Env.MaxIdempotentCallAttempts
var doOnce sync.Once
var NetClient fasthttp.Client

func Init() {
	doOnce.Do(func() {
		NetClient = fasthttp.Client{
			MaxConnsPerHost:           *MaxConnPerHost,
			MaxConnWaitTimeout:        time.Duration(*MaxConnWaitTimeout) * time.Second,
			ReadTimeout:               time.Duration(*ReadTimeout) * time.Second,
			MaxIdemponentCallAttempts: *MaxIdempotentCallAttempts,
			MaxIdleConnDuration:       5 * time.Second,
			MaxConnDuration:           5 * time.Second,
		}
	})
}

func SendGetRequest(url string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	res := fasthttp.AcquireResponse()
	err := NetClient.Do(req, res)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	return res.StatusCode(), res.Body(), err
}
