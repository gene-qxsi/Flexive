package models

import "time"

type Comment struct {
	ID        int       `gorm:"type:int;primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    int       `gorm:"type:bigint;not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	PostID    int       `gorm:"type:bigint;not null;index" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:onDelete:CASCADE" json:"post"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
}
