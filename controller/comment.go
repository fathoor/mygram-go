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

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	Comment := model.Comment{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	Comment.UserId = userId

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

func CommentGetAll(c *gin.Context) {
	db := database.GetDB()
	var (
		Comments []model.Comment
		Data     []interface{}
	)

	err := db.Debug().Preload("User").Preload("Photo").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	for i := range Comments {
		User := gin.H{
			"id":       Comments[i].User.ID,
			"email":    Comments[i].User.Email,
			"username": Comments[i].User.Username,
		}

		Photo := gin.H{
			"id":        Comments[i].Photo.ID,
			"title":     Comments[i].Photo.Title,
			"caption":   Comments[i].Photo.Caption,
			"photo_url": Comments[i].Photo.PhotoUrl,
			"user_id":   Comments[i].Photo.UserId,
		}

		Data = append(Data, gin.H{
			"id":         Comments[i].ID,
			"message":    Comments[i].Message,
			"photo_id":   Comments[i].PhotoId,
			"user_id":    Comments[i].UserId,
			"updated_at": Comments[i].UpdatedAt,
			"created_at": Comments[i].CreatedAt,
			"User":       User,
			"Photo":      Photo,
		})
	}

	c.JSON(http.StatusOK, Data)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := model.Comment{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	Comment.ID = uint(commentId)
	Comment.UserId = userId

	err := db.Debug().Model(&Comment).Preload("User").Preload("Photo").Where("id = ?", commentId).Updates(model.Comment{Message: Comment.Message}).Select("photo_id").Scan(&Comment.PhotoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	userId := int(auth["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := model.Comment{}

	Comment.ID = uint(commentId)
	Comment.UserId = userId

	err := db.Debug().Where("id = ? AND user_id = ?", commentId, userId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
