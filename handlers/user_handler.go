package handlers

import (
	services_get "back/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services_get.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services_get.NewUserService(),
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers2()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}
