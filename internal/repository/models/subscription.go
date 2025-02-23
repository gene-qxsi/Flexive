package models

import "time"

type Subscription struct {
	UserID    int       `gorm:"primaryKey;type:int;not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ChannelID int       `gorm:"primaryKey;type:int;not null"`
	Channel   Channel   `gorm:"foreignKey:ChannelID;constraint:OnDelete:CASCADE"`
	Role      string    `gorm:"type:varchar(10);not null;default:USER"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
