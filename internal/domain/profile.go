package domain

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type Profile struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Bio       string     `json:"bio,omitempty"`
	AvatarURL string     `json:"avatar_url,omitempty"`
	Website   string     `json:"website,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
}

func ToORMProfile(dom *Profile) *models.Profile {
	return &models.Profile{
		ID:        dom.ID,
		UserID:    dom.UserID,
		Bio:       dom.Bio,
		AvatarURL: dom.AvatarURL,
		Website:   dom.Website,
		Birthday:  dom.Birthday,
	}
}

func ToDomainProfile(orm *models.Profile) *Profile {
	return &Profile{
		ID:        orm.ID,
		UserID:    orm.UserID,
		Bio:       orm.Bio,
		AvatarURL: orm.AvatarURL,
		Website:   orm.Website,
		Birthday:  orm.Birthday,
	}
}
