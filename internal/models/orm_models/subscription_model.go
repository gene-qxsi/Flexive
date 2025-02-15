package orm_models

import "time"

type Subscription struct {
	UserID    int       `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	ChannelID int       `gorm:"not null;index" json:"channel_id"`
	Channel   Channel   `gorm:"forignKey:ChannelID;constraint:onDelete:CASCADE" json:"channel"`
	CreatedAt time.Time `gorm:"default:now()" json:"create_at"`
}
