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
	handler := NewStandardServerHandler()
	ResetServerHandler(handler)
	if err := Run(ctx, 8080); err != nil {
		t.Fatal(err)
	}
}
func TestDaemon(t *testing.T) {
	handler := NewStandardServerHandler()
	ResetServerHandler(handler)
	Daemon(ctx, 8080)
	defer Close(ctx)

	t.Log("hello world")

	time.Sleep(time.Second * 5)
}

func TestHandler(t *testing.T) {
	logger.ResetLoggerWithOptions(logger.WithCallerHook())
	handler := NewServerHandlerWithOptions(
		WithLoggingRequest(false),
		WithLoggingResponse(false),
	)
	ResetServerHandler(handler)

	handler.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	handler.Any("hello/:id", NewHandlerFunc(HandleHello))
	handler.GET("hi", NewHandlerFuncEnhanced(HandleHi))

	if err := Run(ctx, 8080); err != nil {
		t.Fatal(err)
	}
}

type req struct {
	Id    string `json:"id" form:"id" uri:"id"`
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Token string `json:"token" header:"token-aa"`
}

type resp struct {
	SayHi string `json:"say_hi"`
}

func HandleHello(ctx context.Context, req *req) (*resp, error) {
	logger.Infof(ctx, "%+v", req)
	return &resp{
		SayHi: "hello world!" + req.Name,
	}, ErrorWithCode(CodeBadRequest)
}
func HandleHi(c *gin.Context, req *req) (*string, error) {
	ctx := c.Request.Context()
	logger.Infof(ctx, "%+v", req)
	c.JSON(http.StatusInternalServerError, "overwrite the default body struct")
	msg := "hello world!"
	return &msg, errors.New("test err")
}
