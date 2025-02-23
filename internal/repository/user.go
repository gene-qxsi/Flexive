package repository

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	const op = "internal/api/repositories/user_repo.go/CreateUser()"
	err := u.db.Debug().Create(user).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}
	fmt.Println(user.PasswordHash)

	return user, nil
}

func (u *UserRepo) GetUser(id int) (*models.User, error) {
	const op = "internal/api/repositories/user_repo.go/GetUser()"

	var user models.User
	err := u.db.Debug().
		Preload("Channels", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, title, user_id, created_at")
		}).
		First(&user, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &user, nil
}

func (u *UserRepo) GetUsers() ([]models.User, error) {
	const op = "internal/api/repositories/user_repo.go/GetUsers()"

	var users []models.User
	err := u.db.Debug().Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return users, nil
}

func (u *UserRepo) UpdateUser(id int, values map[string]interface{}) (*models.User, error) {
	const op = "internal/api/repositories/user_repo.go/UpdateUser()"

	result := u.db.Debug().Model(&models.User{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", fmt.Sprintf("user with ID %d not found or no changes made", id), op)
	}

	var updatedUser models.User
	err := u.db.Debug().First(&updatedUser, id).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &updatedUser, nil
}

func (u *UserRepo) DeleteUser(id int) error {
	const op = "internal/api/repositories/user_repo.go/DeleteUser()"

	result := u.db.Debug().Delete(&models.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", result.Error.Error(), op)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-2: %s. ПУТЬ: %s", fmt.Sprintf("user with ID %d not found", id), op)
	}

	return nil
}

func (u *UserRepo) FindByUsername(username string) (*models.User, error) {
	const op = "internal/repository/user_repo.go/FindByUsername()"

	var user models.User
	err := u.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &user, nil
}

func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	const op = "internal/repository/user_repo.go/FindByEmail()"

	var user models.User
	err := u.db.Debug().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &user, nil
}

func (u *UserRepo) FindByEmailAndPasword(email, password string) (*models.User, error) {
	const op = "internal/repository/user_repo.go/FindByEmailAndPasword()"

	var user models.User
	err := u.db.Debug().Where("email = ?", email).Where("password = ?", password).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &user, nil
}
