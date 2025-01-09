package services_get

import (
	"back/models"
)

type UserService struct {
	// Add database connection or other dependencies here
	// just any example code

	// db *sql.DB
	// cache *redis.Client
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers2() ([]models.User, error) {
	// Implement database logic here
	return []models.User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
	}, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	// Implement database logic here
	return &models.User{
		ID:    id,
		Name:  "John Doe",
		Email: "johnydoe@gmai.com",
	}, nil
}
