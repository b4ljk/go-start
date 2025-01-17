// router/router.go
package router

import (
	"back/router/admin"
	"back/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	baseController, err := services.NewController()
	if err != nil {
		panic(err)
	}

	admin.AdminController{Controller: *baseController}.Register(r.Group("/admin"))

	return r
}

func setupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// User routes
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}
	}
}
