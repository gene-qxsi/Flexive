package models

import "time"

type Profile struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	UserID    int        `gorm:"uniqueIndex;not null" json:"user_id"`
	Bio       string     `gorm:"type:text" json:"bio,omitempty"`
	AvatarURL string     `gorm:"type:varchar(255)" json:"avatar_url,omitempty"`
	Website   string     `gorm:"type:varchar(128)" json:"website,omitempty"`
	Birthday  *time.Time `gorm:"type:date" json:"birthday,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
