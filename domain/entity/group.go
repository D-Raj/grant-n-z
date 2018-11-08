package entity

import (
	"time"
)

type Group struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Domain    string    `gorm:"type:varchar(128);not null" json:"domain" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e Group) GetTableName() string {
	return "groups"
}
