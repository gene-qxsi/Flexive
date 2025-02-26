package models

import "time"

type Message struct {
	ID        int       `gorm:"type:int;primaryKey"`
	ChatID    int       `gorm:"type:int;not null;index"`
	Chat      Chat      `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	UserID    int       `gorm:"type:int;not null;index"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
