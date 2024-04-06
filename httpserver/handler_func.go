package httpserver

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/logger"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newBodyWithError(err *Error) *Body {
	return &Body{
		Code:    int(err.Code),
		Message: err.Error(),
	}
}

type HandlerFn[RequestT any, ResponseT any] func(context.Context, *RequestT, *ResponseT) error
type HandlerFnEnhanced[RequestT any, ResponseT any] func(*gin.Context, *RequestT, *ResponseT) error

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
		requestPtr := reflect.New(reflect.TypeOf((*RequestT)(nil)).Elem()).Interface()
		responsePtr := reflect.New(reflect.TypeOf((*ResponseT)(nil)).Elem()).Interface()

		if err := c.ShouldBind(requestPtr); err != nil {
			logger.WithError(err).Errorf(ctx, "failed to bind, request: %+v, response: %+v", requestPtr, responsePtr)
			c.PureJSON(http.StatusBadRequest, newBodyWithError(ErrorWithBadRequest()))
		}

		var err error
		if ginCtx {
			err = fnEnhanced(c, requestPtr.(*RequestT), responsePtr.(*ResponseT))
		} else {
			err = fn(ctx, requestPtr.(*RequestT), responsePtr.(*ResponseT))
		}
		if c.Writer.Written() {
			return
		}
		if err != nil {
			c.Writer.Flush()
			var body *Body
			status := http.StatusOK
			switch e := err.(type) {
			case *Error:
				body = newBodyWithError(e)
				if e.Status != 0 {
					status = e.Status
				}
			case error:
				body = &Body{
					Code:    int(CodeInternalServerError),
					Message: e.Error(),
				}
			}
			logger.WithError(err).Errorf(ctx, "failed to handle, request: %+v, response: %+v", requestPtr, responsePtr)
			c.PureJSON(status, body)
			return
		}
		c.PureJSON(http.StatusOK, &Body{
			Code: int(CodeOK),
			Data: responsePtr,
		})
	}
}
