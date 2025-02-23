package dto

import (
	"time"
)

type ReactionDTO struct {
	UserID    int       `json:"user_id,omitempty"`
	User      UserDTO   `json:"user,omitempty"`
	PostID    int       `json:"post_id,omitempty"`
	Post      PostDTO   `json:"post,omitempty"`
	Reaction  string    `json:"reaction,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
