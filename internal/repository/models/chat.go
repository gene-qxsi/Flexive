package models

import "time"

type Chat struct {
	ID        int       `gorm:"type:int;primaryKey"`
	Title     string    `gorm:"type:varchar(128)"`
	UserID    int       `gorm:"type:int;not null;index"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
