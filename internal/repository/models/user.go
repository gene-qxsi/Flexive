package models

import "time"

type User struct {
	ID            int            `gorm:"primaryKey" json:"id"`
	Username      string         `gorm:"type:varchar(64);not null;unique" json:"username"`
	Description   string         `gorm:"type:varchar(256);default:null" json:"description"`
	Email         string         `gorm:"type:varchar(128);not null;unique" json:"email"`
	Password      string         `gorm:"type:varchar(255);not null" json:"password"`
	Role          string         `gorm:"type:varchar(5);not null" json:"role"`
	Birthday      *time.Time     `gorm:"type:timestamp;default:null" json:"birthday"`
	Channels      []Channel      `gorm:"foreignKey:UserID;references:ID" json:"channels"`
	Comments      []Comment      `gorm:"foreignKey:UserID;references:ID" json:"comments"`
	Posts         []Post         `gorm:"foreignKey:UserID;references:ID" json:"posts"`
	Reactions     []Reaction     `gorm:"foreignKey:UserID;references:ID" json:"reactions"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID;references:ID" json:"subscriptions"`
	CreatedAt     time.Time      `gorm:"type:timestamp;default:now()" json:"created_at"`
}
