package httpserver

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/httpserver/middleware"
)

func NewServerHandler() *gin.Engine {
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

func NewStandardServerHandler() *gin.Engine {
	engine := NewServerHandler()
	engine.Use(
		middleware.Recovery(),
		middleware.LoggingRequest(true),
		middleware.LoggingResponse(true),
	)
	return engine
}

func NewServerHandlerWithOptions(opts ...Option) *gin.Engine {
	engine := NewServerHandler()
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

type Option func(*gin.Engine)

type Router interface {
	Init(gin.IRouter)
}

func WithRouters(basePath string, routers ...Router) Option {
	return func(e *gin.Engine) {
		group := e.Group(basePath)
		for _, router := range routers {
			router.Init(group)
		}
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
		e.Use(middleware.LoggingRequest(header))
	}
}

func WithReuseBody() Option {
	return func(e *gin.Engine) {
		e.Use(middleware.ReuseBody())
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
