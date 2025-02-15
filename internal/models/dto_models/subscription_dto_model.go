package dto_models

import (
	"time"
)

type SubscriptionDTO struct {
	UserID    int        `json:"user_id,omitempty"`
	User      UserDTO    `json:"user,omitempty"`
	ChannelID int        `json:"channel_id,omitempty"`
	Channel   ChannelDTO `json:"channel,omitempty"`
	CreatedAt time.Time  `json:"create_at,omitempty"`
}
