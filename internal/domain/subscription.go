package domain

import (
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type Subscription struct {
	UserID  int
	Profile Profile
	// User      User
	ChannelID int
	// Channel   Channel
	Role string
}

func ToORMSubscription(subs Subscription) *models.Subscription {
	return &models.Subscription{
		UserID: subs.UserID,
		// User:      *ToORMUser(&subs.User),
		ChannelID: subs.ChannelID,
		// Channel:   *ToORMChannel(&subs.Channel),
		Role: subs.Role,
	}
}
