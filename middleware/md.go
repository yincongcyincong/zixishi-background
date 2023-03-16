package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			if value == "session-886" {
				c.Next()
				return
			}
		}
		if url := c.Request.URL.String(); url == "/login" || url == "/do_login" || url == "/login/captcha" ||
			strings.Contains(url, "/api/") {
			c.Next()
			return
		}

		if c.ContentType() != "application/json" && c.Request.Method == http.MethodGet {
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}

		c.Abort()
		return
	}
}
