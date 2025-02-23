package domain

import (
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type User struct {
	ID           int     `json:"id"`
	Profile      Profile `json:"profile"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"-"`
}

func ToORMUser(user *User) *models.User {
	return &models.User{
		ID:           user.ID,
		Profile:      *ToORMProfile(&user.Profile),
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func ToDomainUser(orm *models.User) *User {
	return &User{
		ID:           orm.ID,
		Profile:      *ToDomainProfile(&orm.Profile),
		Email:        orm.Email,
		PasswordHash: orm.PasswordHash,
	}
}
