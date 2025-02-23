package models

import "time"

type Subscription struct {
	UserID    int       `gorm:"type:bigint;not null;index"`
	User      User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE"`
	ChannelID int       `gorm:"type:bigint;not null;index"`
	Channel   Channel   `gorm:"forignKey:ChannelID;constraint:onDelete:CASCADE"`
	Role      string    `gorm:"type:varchar(9);not null;default:USER"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()"`
}
