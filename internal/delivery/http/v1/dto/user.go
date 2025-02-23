package dto

import (
	"time"
)

type UserDTO struct {
	ID            int               `json:"id"`
	Username      *string           `json:"username"`
	Email         string            `json:"email"`
	Password      string            `json:"-"`
	Channels      []ChannelDTO      `json:"channels,omitempty"`
	Comments      []CommentDTO      `json:"comments,omitempty"`
	Posts         []PostDTO         `json:"posts,omitempty"`
	Reactions     []ReactionDTO     `json:"reactions,omitempty"`
	Subscriptions []SubscriptionDTO `json:"subscriptions,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
}
