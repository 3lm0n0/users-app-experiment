package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID 		`json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string			`json:"name" gorm:"type:varchar(255);not null"`
	Lastname string 	`json:"lastname" gorm:"type:varchar(255);not null"`
	Email string 		`json:"email" gorm:"type:varchar(255);uniqueIndex;not null" validate:"required"`
	Password string 	`json:"password" gorm:"type:varchar(255)" validate:"required"`
	Role string 		`json:"role" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:NULL"`
	DeletedAt time.Time `gorm:"default:NULL"`
}