package httpserver

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/httpserver/middleware"
)

func NewHandler() *gin.Engine {
	gin.DefaultWriter = io.Discard
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(middleware.Recovery())

	//默认的健康检查接口
	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return engine
}

func NewStandardHandler() *gin.Engine {
	engine := NewHandler()
	engine.Use(
		middleware.Recovery(),
		middleware.ReuseBody(),
		middleware.LoggingRequest(true),
		middleware.LoggingResponse(true),
	)
	return engine
}

func NewHandlerWithOptions(opts ...Option) *gin.Engine {
	engine := NewHandler()
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

type Option func(*gin.Engine)

func WithRouter(router func(*gin.Engine)) Option {
	return func(e *gin.Engine) {
		router(e)
	}
}

func WithMiddlewares(middlewares ...gin.HandlerFunc) Option {
	return func(e *gin.Engine) {
		for _, m := range middlewares {
			e.Use(m)
		}
	}
}

func WithLoggingRequest(header bool) Option {
	return func(e *gin.Engine) {
		e.Use(middleware.ReuseBody(), middleware.LoggingRequest(header))
	}
}
func WithLoggingResponse(header bool) Option {
	return func(e *gin.Engine) {
		e.Use(middleware.LoggingResponse(header))
	}
}

func WithServiceName(name string) Option {
	return func(engine *gin.Engine) {
		engine.Use(func(c *gin.Context) {
			c.Set("service", name)
		})
	}
}
