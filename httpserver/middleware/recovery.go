package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Print error and stack information
				//logger.Errorf("panic: %v", err)
				//logger.Errorf("stack: %s", debug.Stack())
				// Terminate subsequent interface calls
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}

}
