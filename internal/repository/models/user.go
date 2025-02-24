package models

import "time"

type User struct {
	ID            int            `gorm:"primaryKey"`
	Email         string         `gorm:"type:varchar(255);not null;unique"`
	PasswordHash  string         `gorm:"column:password_hash;type:text;not null"`
	Profile       Profile        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Channels      []Channel      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Comments      []Comment      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Posts         []Post         `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt     time.Time      `gorm:"autoUpdateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
}
