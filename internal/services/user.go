package services

import (
	"context"
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/repository"
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
	user.PasswordHash, err = s.Hasher.Hash(user.PasswordHash)
	fmt.Println(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	userORM := domain.ToORMUser(user)
	fmt.Println(userORM.PasswordHash)
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

func (s *UserService) GetUsers() ([]domain.User, error) {
	users, err := s.Repo.GetUsers()
	if err != nil {
		return nil, err
	}
	var usersDomain []domain.User
	for _, user := range users {
		usersDomain = append(usersDomain, *domain.ToDomainUser(&user))

	}
	return usersDomain, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(id int, values map[string]interface{}) (*domain.User, error) {
	userORM, err := s.Repo.UpdateUser(id, values)
	if err != nil {
		return nil, err
	}

	user := domain.ToDomainUser(userORM)
	return user, err
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	userORM, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	user := domain.ToDomainUser(userORM)
	return user, err
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	userORM, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	user := domain.ToDomainUser(userORM)
	return user, err
}

func (s *UserService) FindByEmailAndPasword(ctx context.Context, email, password string) (*domain.User, error) {
	userORM, err := s.Repo.FindByEmailAndPasword(email, password)
	if err != nil {
		return nil, err
	}

	user := domain.ToDomainUser(userORM)
	return user, err
}
