package models

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    int       `gorm:"type:int;not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	PostID    int       `gorm:"type:int;not null;index" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"post"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
