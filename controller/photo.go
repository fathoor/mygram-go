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
		User := gin.H{
			"email":    Photos[i].User.Email,
			"username": Photos[i].User.Username,
		}

		Data = append(Data, gin.H{
			"id":         Photos[i].ID,
			"title":      Photos[i].Title,
			"caption":    Photos[i].Caption,
			"photo_url":  Photos[i].PhotoUrl,
			"user_id":    Photos[i].UserId,
			"created_at": Photos[i].CreatedAt,
			"updated_at": Photos[i].UpdatedAt,
			"User":       User,
		})
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

	Photo.ID = uint(photoId)
	Photo.UserId = userId

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

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := model.Photo{}

	err := db.Debug().Where("id = ? AND user_id = ?", photoId, userId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
