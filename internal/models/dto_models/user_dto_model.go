package dto_models

import (
	"time"
)

type UserDTO struct {
	ID            int               `json:"id,omitempty"`
	Username      string            `json:"username,omitempty"`
	Description   string            `json:"description,omitempty"`
	Email         string            `json:"email,omitempty"`
	Password      string            `json:"-"`
	Birthday      *time.Time        `json:"birthday,omitempty"`
	Channels      []ChannelDTO      `json:"channels,omitempty"`
	Comments      []CommentDTO      `json:"comments,omitempty"`
	Posts         []PostDTO         `json:"posts,omitempty"`
	Reactions     []ReactionDTO     `json:"reactions,omitempty"`
	Subscriptions []SubscriptionDTO `json:"subscriptions,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
}
