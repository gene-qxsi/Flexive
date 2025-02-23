package dto

import "time"

type Profile struct {
	UserID    int        `json:"user_id"`
	Username  string     `json:"username"`
	Bio       string     `json:"bio"`
	AvatarURL string     `json:"avatar_url"`
	Website   *string    `json:"website"`
	Birthday  *time.Time `json:"birthday"`
	Role      *string    `json:"role"`
}
