package controller

import (
	"net/http"

	"github.com/fathoor/mygram-go/database"
	"github.com/fathoor/mygram-go/helper"
	"github.com/fathoor/mygram-go/model"
	"github.com/gin-gonic/gin"
)

var (
	APP_JSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType == APP_JSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType == APP_JSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "Email or password is wrong",
		})
		return
	}

	if !helper.ComparePassword([]byte(User.Password), []byte(password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "Wrong password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
