package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *entity.ErrorResponse)

	InsertUser(user *entity.User) (*entity.User, *entity.ErrorResponse)
}