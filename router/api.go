package router

import (
	"github.com/farrelahmady/my-gram-hacktiv8/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", controllers.Login)
		authRouter.POST("/register", controllers.Register)
	}

	return router
}
