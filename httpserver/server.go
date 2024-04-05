package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ihezebin/oneness/logger"
	"github.com/pkg/errors"
)

var server = &http.Server{}
var closeSecure = true

func Server() *http.Server {
	return server
}

func ResetCloseSecure(secure bool) {
	closeSecure = secure
}

func ResetServerHandler(handler http.Handler) {
	server.Handler = handler
}

func Daemon(ctx context.Context, port uint) {
	go func() {
		if err := Run(ctx, port); err != nil {
			logger.WithError(err).Error(ctx, "daemon http server run err")
		}
	}()
}

func Run(ctx context.Context, port uint) error {
	server.Addr = fmt.Sprintf(":%d", port)
	ch := make(chan os.Signal)
	go func() {
		logger.Infof(ctx, "http server is starting in port: %d", port)
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.WithError(err).Error(ctx, "http server ListenAndServe err")
				ch <- SIGRUNERR
			}
			closeOperate := "closed"
			if closeSecure {
				closeOperate = "shutdown"
			}
			logger.Infof(ctx, "http server %s", closeOperate)
		}
	}()
	// handle signal, to elegant closing server
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, SIGRUNERR)
	logger.Infof(ctx, "got signal %v, http server exit", <-ch)

	if closeSecure {
		return server.Shutdown(ctx)
	}

	return server.Close()
}

func Close(ctx context.Context) error {
	if closeSecure {
		return server.Shutdown(ctx)
	}
	return server.Close()
}
