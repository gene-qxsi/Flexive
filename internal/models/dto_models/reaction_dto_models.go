package dto_models

import "time"

type ReactionDTO struct {
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Reaction  string    `json:"reaction"`
	CreatedAt time.Time `json:"created_at"`
}
