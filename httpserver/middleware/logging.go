package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihezebin/oneness/logger"
)

const maxBodyLen = 1024

func LoggingRequest(header bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		fields := map[string]interface{}{
			"method": c.Request.Method,
			"uri":    c.Request.URL.RequestURI(),
			"remote": c.Request.RemoteAddr,
			"body":   requestBody(c.Request.Body),
		}
		if header {
			fields["header"] = c.Request.Header
		}
		logger.WithFields(fields).Info(ctx, "incoming http request")
		c.Next()
	}
}

func requestBody(body io.ReadCloser) string {
	if body == nil || body == http.NoBody {
		return ""
	}
	bodyData, err := io.ReadAll(body)
	if err != nil {
		return fmt.Sprintf("read request body err: %s", err.Error())
	}
	bodySize := len(bodyData)
	if bodySize > maxBodyLen {
		bodySize = maxBodyLen
	}
	return string(bodyData[:bodySize])
}

func LoggingResponse(header bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		rw := responseWriter{Body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = rw
		c.Next()

		fields := map[string]interface{}{
			"status": fmt.Sprintf("%v %s", c.Writer.Status(), http.StatusText(c.Writer.Status())),
			"body":   responseBody(&rw),
		}
		if header {
			fields["header"] = c.Writer.Header()
		}

		logger.WithFields(fields).Info(ctx, "outgoing http response")
	}
}

func responseBody(rw *responseWriter) string {
	body := rw.Body.Bytes()
	bodyLen := len(body)
	if bodyLen > maxBodyLen {
		bodyLen = maxBodyLen
	}
	return string(body[:bodyLen])
}

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w responseWriter) Write(body []byte) (int, error) {
	// store body
	w.Body.Write(body)
	// write
	return w.ResponseWriter.Write(body)
}
