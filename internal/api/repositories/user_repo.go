package repositories

import (
	"fmt"

	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/internal/storage"
)

type UserRepo struct {
	storage *storage.Storage
}

func NewUserRepo(storage *storage.Storage) *UserRepo {
	return &UserRepo{
		storage: storage,
	}
}

func (u *UserRepo) CreateUser(user *models.User) (int, error) {
	const op = "internal/api/repositories/user_repo.go/CreateUser()"

	err := u.storage.Sdb.Create(user).Error
	if err != nil {
		return 0, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return user.ID, nil
}

func (u *UserRepo) GetUser(id int) (*models.User, error) {
	const op = "internal/api/repositories/user_repo.go/GetUser()"

	var user models.User
	err := u.storage.Sdb.Preload("Channels.User").Preload("Comments.User").Preload("Posts.User").
		Preload("Reactions.User").Preload("Subscriptions.User").First(&user, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &user, nil
}

func (u *UserRepo) GetUsers() ([]models.User, error) {
	const op = "internal/api/repositories/user_repo.go/GetUsers()"

	var users []models.User
	err := u.storage.Sdb.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return users, nil
}

func (u *UserRepo) UpdateUser(id int, values map[string]interface{}) error {
	const op = "internal/api/repositories/user_repo.go/UpdateUser()"

	result := u.storage.Sdb.Model(&models.User{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("user with ID %d not found or no changes made", id), op)
	}

	return nil
}

func (u *UserRepo) DeleteUser(id int) error {
	const op = "internal/api/repositories/user_repo.go/DeleteUser()"

	result := u.storage.Sdb.Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("user with ID %d not found", id), op)
	}

	return nil
}
