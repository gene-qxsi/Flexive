package domain

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type User struct {
	ID          int        `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	Description string     `json:"description"`
	Role        string     `json:"role"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

func ToORMUser(dom *User) *models.User {
	return &models.User{
		ID:          dom.ID,
		Username:    dom.Username,
		Email:       dom.Email,
		Password:    dom.Password,
		Description: dom.Description,
		Role:        dom.Role,
		Birthday:    dom.Birthday,
		CreatedAt:   dom.CreatedAt,
	}
}

func ToDomainUser(orm *models.User) *User {
	return &User{
		ID:          orm.ID,
		Username:    orm.Username,
		Email:       orm.Email,
		Password:    orm.Password,
		Description: orm.Description,
		Role:        orm.Role,
		Birthday:    orm.Birthday,
		CreatedAt:   orm.CreatedAt,
	}
}
