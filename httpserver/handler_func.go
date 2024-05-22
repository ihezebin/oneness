package httpserver

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/logger"
	"github.com/pkg/errors"
)

type Body struct {
	status  int
	Code    Code        `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (b *Body) WithError(err error) *Body {
	b.Code = CodeInternalServerError
	b.Message = err.Error()
	return b
}

func (b *Body) WithErrorx(errx *Error) *Body {
	b.Code = errx.Code
	b.Message = errx.Error()
	if errx.Status != 0 {
		b.status = errx.Status
	}
	return b
}

type HandlerFn[RequestT any, ResponseT any] func(context.Context, *RequestT) (*ResponseT, error)
type HandlerFnEnhanced[RequestT any, ResponseT any] func(*gin.Context, *RequestT) (*ResponseT, error)

func NewHandlerFunc[RequestT any, ResponseT any](fn HandlerFn[RequestT, ResponseT]) gin.HandlerFunc {
	return newHandlerFunc(false, fn, nil)
}

func NewHandlerFuncEnhanced[RequestT any, ResponseT any](fn HandlerFnEnhanced[RequestT, ResponseT]) gin.HandlerFunc {
	return newHandlerFunc(true, nil, fn)
}

func newHandlerFunc[RequestT any, ResponseT any](ginCtx bool, fn HandlerFn[RequestT, ResponseT], fnEnhanced HandlerFnEnhanced[RequestT,
	ResponseT]) gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var err error
		var responsePtr *ResponseT

		body := &Body{
			status: http.StatusOK,
			Code:   CodeOK,
			Data:   responsePtr,
		}

		requestPtr := reflect.New(reflect.TypeOf((*RequestT)(nil)).Elem()).Interface()

		if err = c.ShouldBind(requestPtr); err != nil {
			logger.WithError(err).Errorf(ctx, "failed to bind, uri: %s", c.Request.RequestURI)
			body = body.WithErrorx(ErrorWithBadRequest())
			c.PureJSON(http.StatusBadRequest, body)
			return
		}

		if len(c.Params) > 0 {
			if err = c.ShouldBindUri(requestPtr); err != nil {
				logger.WithError(err).Errorf(ctx, "failed to bind uri, uri: %s, param: %+v", c.Request.RequestURI, c.Params)
				body = body.WithErrorx(ErrorWithBadRequest())
				c.PureJSON(http.StatusBadRequest, body)
				return
			}
		}

		if ginCtx {
			responsePtr, err = fnEnhanced(c, requestPtr.(*RequestT))
			if c.Writer.Written() {
				return
			}
		} else {
			responsePtr, err = fn(ctx, requestPtr.(*RequestT))
		}

		// handle error
		if err != nil {
			var errx *Error
			if errors.As(err, &errx) {
				body = body.WithErrorx(errx)
			} else {
				body = body.WithError(err)
			}
			body.Data = nil
		} else {
			body.Data = responsePtr
		}

		// handle success
		c.PureJSON(body.status, body)
	}
}
