package models

import "time"

type Channel struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(128);not null" json:"title"`
	Description *string   `gorm:"type:varchar(1024);" json:"description,omitempty"`
	UserID      int       `gorm:"type:int;not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
