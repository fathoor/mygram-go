package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/fathoor/mygram-go/database"
	"github.com/fathoor/mygram-go/helper"
	"github.com/fathoor/mygram-go/model"
	"github.com/gin-gonic/gin"
)

var APP_JSON = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	User := model.User{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"msg":   "Age must be number",
				})
				return
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"msg":   "Age must be number",
				})
				return
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Email already exists",
			})
			return
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Username already exists",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   err.Error(),
			})
		}
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
	User := model.User{}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"msg":   "Email is not registered",
			})
			return
		}
	} else if !helper.ComparePassword([]byte(User.Password), []byte(password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "Password is incorrect",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	User := model.User{}
	UserId := uint(auth["id"].(float64))

	userId, e := strconv.Atoi(c.Param("userId"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   "Invalid user id",
		})
		return
	}

	if contentType == APP_JSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	User.ID = UserId

	if userId != int(UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "You can't update other user",
		})
		return
	}

	err := db.Debug().Model(&User).Where("id = ?", userId).Updates(&User).Select("age").Scan(&User.Age).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Email already exists",
			})
			return
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Username already exists",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	User := model.User{}
	UserId := uint(auth["id"].(float64))

	err := db.Debug().Model(&User).Where("id = ?", UserId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
