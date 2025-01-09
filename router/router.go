// router/router.go
package router

import (
	"back/handlers"
	"back/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	startTime := time.Now()

	// Initialize services
	log.Printf("STARTIIING")

	serviceStartTime := time.Now()
	userService, err := services.NewUserService()
	if err != nil {
		log.Fatalf("Failed to initialize user service: %v", err)
	}
	log.Printf("Services initialization took: %v", time.Since(serviceStartTime))

	// Initialize handlers
	handlerStartTime := time.Now()
	userHandler := handlers.NewUserHandler(userService)
	log.Printf("Handlers initialization took: %v", time.Since(handlerStartTime))

	// Setup Gin
	r := gin.Default()

	// Setup routes
	routesStartTime := time.Now()
	setupRoutes(r, userHandler)
	log.Printf("Routes setup took: %v", time.Since(routesStartTime))

	log.Printf("Total router setup took: %v", time.Since(startTime))
	return r
}

func setupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// User routes
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetUsers)
		}
	}
}
