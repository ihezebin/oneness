package httpclient

import "github.com/go-resty/resty/v2"

func transmitDataThroughHeaderMiddleware(c *resty.Client, request *resty.Request) error {
	ctx := request.Context()
	// get and do something from ctx
	_ = ctx
	//request.Header.Set(k, v)
	return nil
}
