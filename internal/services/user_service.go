package services

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/pkg/mappers"
)

type UserService struct {
	Repo   *repository.UserRepo
	Hasher *BcryptHasher
}

func NewUserService(repo *repository.UserRepo, hasher *BcryptHasher) *UserService {
	return &UserService{Repo: repo, Hasher: hasher}
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	var err error
	user.Password, err = s.Hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	userORM := domain.ToORMUser(user)
	userORM, err = s.Repo.CreateUser(userORM)
	if err != nil {
		return nil, err
	}

	user = domain.ToDomainUser(userORM)
	return user, err
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	userRepo, err := s.Repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	user := domain.ToDomainUser(userRepo)
	return user, err
}

func (s *UserService) GetUsers() ([]dto.UserDTO, error) {
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

func (s *UserService) UpdateUser(id int, values map[string]interface{}) (*dto.UserDTO, error) {
	user, err := s.Repo.UpdateUser(id, values)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*dto.UserDTO, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}

func (s *UserService) FindByEmailAndPasword(ctx context.Context, email, password string) (*dto.UserDTO, error) {
	user, err := s.Repo.FindByEmailAndPasword(email, password)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.UserToDTO(*user)
	return &userDTO, err
}
