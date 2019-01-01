package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Uuid        uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	Username    string    `gorm:"type:varchar(128)"json:"username"`
	DisplayName string    `gorm:"type:varchar(128)"json:"display_name"`
	Email       string    `gorm:"type:varchar(128);not null;index:email"json:"email"`
	Password    string    `gorm:"type:varchar(128);not null"json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserReq struct {
	Email       string `validate:"required,email"json:"email"`
	Password    string `validate:"required"json:"password"`
}

func (e User) TableName() string {
	return "users"
}
