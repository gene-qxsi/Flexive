package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Views     int       `gorm:"type:int;not null;default:0" json:"views"`
	UserID    int       `gorm:"type:int;not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	ChannelID int       `gorm:"type:int;not null;index" json:"channel_id"`
	Channel   Channel   `gorm:"foreignKey:ChannelID;constraint:OnDelete:CASCADE" json:"channel"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
