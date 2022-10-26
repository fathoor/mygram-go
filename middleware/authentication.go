package middleware

import (
	"net/http"

	"github.com/fathoor/mygram-go/helper"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken, e := helper.ValidateToken(c)
		_ = validateToken

		if e != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthenticated",
				"msg":   e.Error(),
			})
			return
		}

		c.Set("auth", validateToken)
		c.Next()
	}
}
