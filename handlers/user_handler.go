package handlers

import (
	services_get "back/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services_get.UserService
}

func NewUserHandler(user_service *services_get.UserService) *UserHandler {

	return &UserHandler{
		userService: user_service,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.userService.GetUserByID(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}
