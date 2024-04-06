package httpserver

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/logger"
	"github.com/pkg/errors"
)

var ctx = context.Background()

func TestServer(t *testing.T) {
	handler := NewStandardHandler()
	ResetServerHandler(handler)
	if err := Run(ctx, 8080); err != nil {
		t.Fatal(err)
	}
}
func TestDaemon(t *testing.T) {
	handler := NewStandardHandler()
	ResetServerHandler(handler)
	Daemon(ctx, 8080)
	defer Close(ctx)

	t.Log("hello world")

	time.Sleep(time.Second * 5)
}

func TestHandler(t *testing.T) {
	handler := NewHandlerWithOptions(
		WithCorsEveryOrigin(),
		WithLoggingRequest(false),
		WithLoggingResponse(false),
		WithRouter(func(e *gin.Engine) {
			e.GET("ping", func(c *gin.Context) {
				c.String(http.StatusOK, "pong")
			})
			e.GET("hello", NewHandlerFunc(HandleHello))
			e.GET("hi", NewHandlerFuncEnhanced(HandleHi))
		}),
	)
	ResetServerHandler(handler)
	if err := Run(ctx, 8080); err != nil {
		t.Fatal(err)
	}
}

type req struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

type resp struct {
	SayHi string `json:"say_hi"`
}

func HandleHello(ctx context.Context, req *req, resp *resp) error {
	logger.Infof(ctx, "%+v", req)
	resp.SayHi = "hello world!"
	return nil
}
func HandleHi(c *gin.Context, req *req, data *string) error {
	ctx := c.Request.Context()
	logger.Infof(ctx, "%+v", req)
	*data = "hello world!"
	c.JSON(http.StatusInternalServerError, "overwrite the default body struct")
	return errors.New("test err")
}
