package middleware

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

func ReuseBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := bytes.Buffer{}
		c.Request.Body = io.NopCloser(io.TeeReader(c.Request.Body, &buf))
		c.Next()
		c.Request.Body = io.NopCloser(&buf)
	}
}
