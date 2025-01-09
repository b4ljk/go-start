package router

import (
	"back/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group API routes
	api := r.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			userHandler := handlers.NewUserHandler()
			users.GET("/", userHandler.GetUsers)
		}
	}
	// default hello world
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!!",
		})
	})

	return r
}
