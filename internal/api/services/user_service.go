package services

import (
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type UserService struct {
	Repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) (int, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.Repo.GetUser(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Repo.GetUsers()
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(id int, values map[string]interface{}) error {
	return s.Repo.UpdateUser(id, values)
}
