package dto_models

import (
	"time"
)

type PostDTO struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Content     string     `json:"content,omitempty"`
	Views       int        `json:"views,omitempty"`
	UserID      int        `json:"user_id,omitempty"`
	User        UserDTO    `json:"user,omitempty"`
	ChannelID   int        `json:"channel_id,omitempty"`
	Channel     ChannelDTO `json:"channel,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
}
