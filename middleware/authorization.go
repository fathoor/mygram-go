package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/fathoor/mygram-go/database"
	"github.com/fathoor/mygram-go/model"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   "Invalid photo id",
			})
			return
		}

		auth := c.MustGet("auth").(jwt.MapClaims)
		userId := int(auth["id"].(float64))
		Photo := model.Photo{}

		err = db.Debug().Select("user_id").First(&Photo, photoId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"msg":   "Photo not found",
			})
			return
		}

		if Photo.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "You are not authorized to access this photo",
			})
			return
		}

		c.Next()
	}
}