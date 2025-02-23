package domain

import "github.com/gene-qxsi/Flexive/internal/repository/models"

type Channel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}

func ToDomainChannel(orm *models.Channel) *Channel {
	return &Channel{
		ID:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		UserID:      orm.UserID,
	}
}
