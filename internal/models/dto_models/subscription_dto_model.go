package dto_models

import "time"

type SubscriptionDTO struct {
	UserID    int       `json:"user_id"`
	ChannelID int       `json:"channel_id"`
	CreatedAt time.Time `json:"create_at"`
}
