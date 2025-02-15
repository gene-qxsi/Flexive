package orm_models

import "time"

type Channel struct {
	ID          int       `gorm:"type:int;primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(64);size:64;not null" json:"title"`
	Description string    `gorm:"type:varchar(1024);size:1024" json:"description"`
	UserID      int       `gorm:"type:bigint;not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
}
