package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Service struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	Name      string    `gorm:"type:varchar(128);not null"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Service) TableName() string {
	return "services"
}
