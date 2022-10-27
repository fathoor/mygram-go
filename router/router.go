package router

import (
	"github.com/fathoor/mygram-go/controller"
	"github.com/fathoor/mygram-go/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.PUT("/:userId", middleware.Authentication(), controller.UserUpdate)
		userRouter.DELETE("/", middleware.Authentication(), controller.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.PhotoCreate)
		photoRouter.GET("/", controller.PhotoGetAll)
	}

	return r
}
