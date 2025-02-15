package dto_models

import "time"

type UserDTO struct {
	ID          int        `json:"id"`
	Username    string     `json:"username"`
	Description string     `json:"description"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Birthday    *time.Time `json:"birthday"`
	CreatedAt   time.Time  `json:"created_at"`
}
