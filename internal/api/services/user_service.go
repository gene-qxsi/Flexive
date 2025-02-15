package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/models/dto_models"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type UserService struct {
	Repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) (*dto_models.UserDTO, error) {
	user, err := s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}

func (s *UserService) GetUser(id int) (*dto_models.UserDTO, error) {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}

func (s *UserService) GetUsers() ([]dto_models.UserDTO, error) {
	users, err := s.Repo.GetUsers()
	if err != nil {
		return nil, err
	}
	usersDTOs := mappers.UsersToDTOs(users)
	return usersDTOs, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(id int, values map[string]interface{}) (*dto_models.UserDTO, error) {
	user, err := s.Repo.UpdateUser(id, values)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}
