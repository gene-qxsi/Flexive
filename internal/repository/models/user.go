package models

import "time"

type User struct {
	// Username      string         `gorm:"type:varchar(255);not null;unique"`
	ID            int            `gorm:"primaryKey"`
	Email         string         `gorm:"type:varchar(255);not null;unique"`
	PasswordHash  string         `gorm:"column:password_hash;type:text;not null"`
	Channels      []Channel      `gorm:"foreignKey:UserID;references:ID"`
	Comments      []Comment      `gorm:"foreignKey:UserID;references:ID"`
	Posts         []Post         `gorm:"foreignKey:UserID;references:ID"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt     time.Time      `gorm:"type:timestamp;default:now()"`
}
