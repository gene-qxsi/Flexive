package dto_models

import "time"

type ChannelDTO struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	User        UserDTO   `json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
