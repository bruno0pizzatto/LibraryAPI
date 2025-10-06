package routes

import (
	"github.com/bruno0pizzatto/LibraryAPI/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
		}
		users := main.Group("users")
		{
			users.GET(":/id", controllers.ShowUser)
			users.GET("/", controllers.ShowUsers)
			users.POST("/", controllers.CreateUser)
		}
	}

	return router
}
