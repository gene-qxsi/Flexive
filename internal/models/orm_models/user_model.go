package orm_models

import (
	"time"

	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
)

type User struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"type:varchar(64);not null;unique" json:"username"`
	Description string    `gorm:"type:varchar(256)" json:"description"`
	Email       string    `gorm:"type:varchar(128);not null;unique" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	Birthday    time.Time `gorm:"type:TIMESTAMP;default:NULL" json:"birthday"`
	Channels    models.Channel
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
}
