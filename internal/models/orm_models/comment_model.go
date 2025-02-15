package orm_models

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	PostID    int       `gorm:"not null;index" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:onDelete:CASCADE" json:"post"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}
