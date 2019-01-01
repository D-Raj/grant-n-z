package entity

import (
	"time"
)

type Role struct {
	Id         int       `gorm:"primary_key"json:"id"`
	Permission string    `gorm:"type:varchar(128);not null"validate:"required"json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (e Role) TableName() string {
	return "roles"
}
