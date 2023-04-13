package router

import (
	"github.com/farrelahmady/my-gram-hacktiv8/controllers"
	"github.com/farrelahmady/my-gram-hacktiv8/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		// Login
		authRouter.POST("/login", controllers.Login)
		// Register
		authRouter.POST("/register", controllers.Register)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		// GetAllPhotos
		photoRouter.GET("/", controllers.GetAllPhotos)
		// GetOnePhoto
		photoRouter.GET("/:id", controllers.GetPhoto)
		// CreatePhoto
		photoRouter.POST("/", controllers.CreatePhoto)
		// UpdatePhoto
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// DeletePhoto
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		// GetAllPhotos
		commentRouter.GET("/", controllers.GetAllComments)
		// GetOneComment
		commentRouter.GET("/:id", controllers.GetComment)
		// CreateComment
		commentRouter.POST("/", controllers.CreateComment)
		// UpdateComment
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// DeleteComment
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
