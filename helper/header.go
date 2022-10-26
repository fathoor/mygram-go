package helper

import "github.com/gin-gonic/gin"

func GetContentType(c *gin.Context) string {
	return c.ContentType()
}
