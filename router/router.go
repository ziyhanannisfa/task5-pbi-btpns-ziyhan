package router

import (
	"PBI/controllers"
	"PBI/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
		userGroup.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/", middlewares.AuthMiddleware(), controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetPhotos)
		photoGroup.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdatePhotoByID)
		photoGroup.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeletePhotoByID)
	}

	return r
}
