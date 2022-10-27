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

func SocialMediaCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	SocialMedia := model.SocialMedia{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&SocialMedia); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&SocialMedia); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	SocialMedia.UserId = userId

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func SocialMediaGetAll(c *gin.Context) {
	db := database.GetDB()
	var (
		SocialMedias []model.SocialMedia
		Data         []interface{}
	)

	err := db.Debug().Preload("User").Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	for i := range SocialMedias {
		User := gin.H{
			"id":       SocialMedias[i].User.ID,
			"username": SocialMedias[i].User.Username,
			"email":    SocialMedias[i].User.Email,
		}

		Data = append(Data, gin.H{
			"id":               SocialMedias[i].ID,
			"name":             SocialMedias[i].Name,
			"social_media_url": SocialMedias[i].SocialMediaUrl,
			"user_id":          SocialMedias[i].UserId,
			"created_at":       SocialMedias[i].CreatedAt,
			"updated_at":       SocialMedias[i].UpdatedAt,
			"User":             User,
		})
	}

	c.JSON(http.StatusOK, Data)
}

func SocialMediaUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := model.SocialMedia{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&SocialMedia); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&SocialMedia); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	SocialMedia.ID = uint(socialMediaId)
	SocialMedia.UserId = userId

	err := db.Debug().Model(&SocialMedia).Preload("User").Where("id = ?", socialMediaId).Updates(model.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"created_at":       SocialMedia.CreatedAt,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func SocialMediaDelete(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := model.SocialMedia{}

	err := db.Debug().Model(&SocialMedia).Where("id = ? AND user_id = ?", socialMediaId, userId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
