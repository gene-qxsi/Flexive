package domain

import "time"

type Subscription struct {
	UserID    int
	User      User
	ChannelID int
	// Channel   Channel
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
