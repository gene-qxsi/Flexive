package dto

import (
	"time"
)

type SubscriptionDTO struct {
	UserID    int        `json:"user_id,omitempty"`
	User      UserDTO    `json:"user,omitempty"`
	ChannelID int        `json:"channel_id,omitempty"`
	Channel   ChannelDTO `json:"channel,omitempty"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:now()"`
}
