package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsEveryOrigin() gin.HandlerFunc {
	return func(c *gin.Context) {
		Cors(c.GetHeader("Origin"))
	}
}

func Cors(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
		//context.Header("Access-Control-Allow-Origin", "*")
		// 变相的 *, 谁请求就放心谁
		c.Header("Access-Control-Allow-Origin", c.GetHeader(origin))
		// 必须，设置服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
		// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
