package domain

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type Profile struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Username  string     `json:"username"`
	Bio       string     `json:"bio"`
	AvatarURL string     `json:"avatar_url"`
	Website   *string    `json:"website"`
	Birthday  *time.Time `json:"birthday"`
	Role      *string    `json:"role"`
}

func ToORMProfile(dom *Profile) *models.Profile {
	return &models.Profile{
		ID:        dom.ID,
		UserID:    dom.UserID,
		Username:  dom.Username,
		Bio:       dom.Bio,
		AvatarURL: dom.AvatarURL,
		Website:   dom.Website,
		Birthday:  dom.Birthday,
		Role:      dom.Role,
	}
}

func ToDomainProfile(orm *models.Profile) *Profile {
	return &Profile{
		ID:        orm.ID,
		UserID:    orm.UserID,
		Username:  orm.Username,
		Bio:       orm.Bio,
		AvatarURL: orm.AvatarURL,
		Website:   orm.Website,
		Birthday:  orm.Birthday,
		Role:      orm.Role,
	}
}
