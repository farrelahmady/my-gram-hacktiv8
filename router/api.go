package router

import (
	"github.com/farrelahmady/my-gram-hacktiv8/controllers"
	"github.com/farrelahmady/my-gram-hacktiv8/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", controllers.Login)
		authRouter.POST("/register", controllers.Register)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.GET("/:id", controllers.GetPhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}
	return router
}
