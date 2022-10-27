package controller

import (
	"net/http"
	"strconv"

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

func PhotoGetAll(c *gin.Context) {
	db := database.GetDB()
	var (
		Photos []model.Photo
		Data   []interface{}
	)

	err := db.Debug().Preload("User").Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	for i := range Photos {
		User := make(map[string]interface{})
		User["email"] = Photos[i].User.Email
		User["username"] = Photos[i].User.Username

		Photo := make(map[string]interface{})
		Photo["id"] = Photos[i].ID
		Photo["title"] = Photos[i].Title
		Photo["caption"] = Photos[i].Caption
		Photo["photo_url"] = Photos[i].PhotoUrl
		Photo["user_id"] = Photos[i].UserId
		Photo["created_at"] = Photos[i].CreatedAt
		Photo["updated_at"] = Photos[i].UpdatedAt
		Photo["User"] = User

		Data = append(Data, Photo)
	}

	c.JSON(http.StatusOK, Data)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := model.Photo{}

	Photo.ID = uint(photoId)
	Photo.UserId = userId

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

	err := db.Debug().Model(&Photo).Where("id = ? AND user_id = ?", photoId, userId).Updates(model.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"updated_at": Photo.UpdatedAt,
	})
}
