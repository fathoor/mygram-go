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
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.PhotoUpdate)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CommentCreate)
		commentRouter.GET("/", controller.CommentGetAll)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controller.CommentUpdate)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controller.CommentDelete)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controller.SocialMediaCreate)
		socialMediaRouter.GET("/", controller.SocialMediaGetAll)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.SocialMediaUpdate)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.SocialMediaDelete)
	}

	return r
}
