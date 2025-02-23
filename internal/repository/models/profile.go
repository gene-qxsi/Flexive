package models

import "time"

type Profile struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	UserID    int        `gorm:"uniqueIndex;not null" json:"user_id"`
	Username  string     `gorm:"type:varchar(100);not null;unique" json:"username"`
	Bio       string     `gorm:"type:text" json:"bio,omitempty"`
	AvatarURL string     `gorm:"type:text" json:"avatar_url,omitempty"`
	Website   *string    `gorm:"type:text" json:"website,omitempty"`
	Birthday  *time.Time `gorm:"type:date" json:"birthday,omitempty"`
	Role      *string    `gorm:"type:varchar(11);default:USER;not null;check:role IN ('USER','PREMIUM')"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
