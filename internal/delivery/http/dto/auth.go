package dto

import "time"

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// роль при регистрации не указывать, в место ролей сделать платные подписки
type SignUpRequest struct {
	Username string     `json:"username" binding:"required,max=64"`
	Email    string     `json:"email" binding:"required,email,max=128"`
	Password string     `json:"password" binding:"required,max=255"`
	Role     string     `json:"role,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}
