package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fathoor/mygram-go/database"
	"github.com/fathoor/mygram-go/helper"
	"github.com/fathoor/mygram-go/model"
	"github.com/gin-gonic/gin"
)

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	Photo := model.Photo{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	Photo.UserId = userId

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"created_at": Photo.CreatedAt,
	})
}
