package orm_models

import "time"

type Post struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:64;not null" json:"title"`
	Description string    `gorm:"size:256" json:"description"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Views       int       `gorm:"default:0" json:"views"`
	UserID      int       `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	ChannelID   int       `gorm:"not null;index" json:"channel_id"`
	Channel     Channel   `gorm:"foreignKey:ChannelID;constraint:onDelete:CASCADE" json:"channel"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
}
