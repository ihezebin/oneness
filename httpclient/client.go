package httpclient

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

var client = NewClient("")

func Client() *resty.Client {
	return client
}

func ResetBaseHost(baseUrl string) {
	client.SetBaseURL(baseUrl)
}

func ResetClient(c *resty.Client) {
	client = c
}

func NewRequest(ctx context.Context) *resty.Request {
	return client.NewRequest().SetContext(ctx)
}
func NewRequestWithClient(ctx context.Context, c *resty.Client) *resty.Request {
	if c == nil {
		c = client
	}
	return c.OnBeforeRequest(transmitDataThroughHeaderMiddleware).NewRequest().SetContext(ctx)
}

// NewClient 多客户端时使用
func NewClient(host string) *resty.Client {
	return resty.New().
		SetBaseURL(host).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(transmitDataThroughHeaderMiddleware)
}
