package httpserver

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func TestHandlerFunc(t *testing.T) {
	handler := NewHandlerWithOptions(
		WithCorsEveryOrigin(),
		WithLoggingRequest(false),
		WithLoggingResponse(false),
		WithRouter(func(e *gin.Engine) {
			e.GET("ping", func(c *gin.Context) {
				c.String(http.StatusOK, "pong")
			})
		}),
	)
	ResetServerHandler(handler)
	if err := Run(ctx, 8080); err != nil {
		t.Fatal(err)
	}
}
