package interfaces

import (
	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type UserService interface {
	CreateUser(user *models.User) (*dto.UserDTO, error)
	GetUser(id int) (*dto.UserDTO, error)
	GetUsers() ([]dto.UserDTO, error)
	DeleteUser(id int) error
	UpdateUser(id int, values map[string]interface{}) (*dto.UserDTO, error)
}
