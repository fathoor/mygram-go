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

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   "Invalid comment id",
			})
			return
		}

		auth := c.MustGet("auth").(jwt.MapClaims)
		userId := int(auth["id"].(float64))
		Comment := model.Comment{}

		err = db.Debug().Select("user_id").First(&Comment, commentId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"msg":   "Comment not found",
			})
			return
		}

		if Comment.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "You are not authorized to access this comment",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   "Invalid social media id",
			})
			return
		}

		auth := c.MustGet("auth").(jwt.MapClaims)
		userId := int(auth["id"].(float64))
		SocialMedia := model.SocialMedia{}

		err = db.Debug().Select("user_id").First(&SocialMedia, socialMediaId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"msg":   "Social media not found",
			})
			return
		}

		if SocialMedia.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "You are not authorized to access this social media",
			})
			return
		}

		c.Next()
	}
}
