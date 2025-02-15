package orm_models

import "time"

type Channel struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:64;not null" json:"title"`
	Description string    `gorm:"size:1024" json:"description"`
	UserID      int       `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
}
