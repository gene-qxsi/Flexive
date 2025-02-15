package dto_models

import "time"

type PostDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Views       int       `json:"views"`
	UserID      int       `json:"user_id"`
	ChannelID   int       `json:"channel_id"`
	CreatedAt   time.Time `json:"created_at"`
}
