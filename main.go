package main

import (
	"back/router"
)

func main() {
	// Load config

	// Initialize router
	r := router.SetupRouter()

	// Start server
	r.Run(":8080")
}

// router/router.go

// handlers/user_handler.go

// models/user.go

// services/user_service.go
