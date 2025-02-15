package dto_models

import "time"

type CommentDTO struct {
	ID        int       `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	User      UserDTO   `json:"user,omitempty"`
	PostID    int       `json:"post_id,omitempty"`
	Post      PostDTO   `json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
