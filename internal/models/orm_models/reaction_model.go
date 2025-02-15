package orm_models

import "time"

type Reaction struct {
	UserID    int       `gorm:"primaryKey" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	PostID    int       `gorm:"primaryKey" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:onDelete:CASCADE" json:"post"`
	Reaction  string    `gorm:"type:VARCHAR(20);not null" json:"reaction"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}
