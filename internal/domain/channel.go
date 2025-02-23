package domain

import "github.com/gene-qxsi/Flexive/internal/repository/models"

type Channel struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	UserID      int     `json:"user_id"`
	User        User    `json:"user"`
}

func ToDomainChannel(orm *models.Channel) *Channel {
	return &Channel{
		ID:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		UserID:      orm.UserID,
		// User:        *ToDomainUser(&orm.User),
	}
}

func ToORMChannel(channel *Channel) *models.Channel {
	return &models.Channel{
		ID:          channel.ID,
		Title:       channel.Title,
		Description: channel.Description,
		UserID:      channel.UserID,
		// User:        *ToORMUser(&channel.User),
	}
}
