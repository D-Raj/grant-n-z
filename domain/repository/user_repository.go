package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type UserRepository interface {
	FindByEmail(email string) *entity.User

	FindByUuid(uuidStr string) (u *entity.User, err error)

	FindByUserName(username string) (u *entity.User, err error)

	Save(user entity.User) (u *entity.User, err error)

	Update(user entity.User) (u *entity.User, err error)

	UpdateUserColumn(user entity.User, column string) (u *entity.User, err error)
}
