package middleware

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
)

func ReuseBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			data = []byte(fmt.Sprintf("read body error: %v", err))
		}
		c.Request.Body = io.NopCloser(bytes.NewReader(data))
		c.Next()
	}

}
