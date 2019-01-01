package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type UserRepositoryImpl struct {
}

func (ur UserRepositoryImpl) NewUserRepository() UserRepository {
	return ur
}

func (ur UserRepositoryImpl) FindByEmail(email string) *entity.User {
	var user entity.User
	infra.Db.Where("email = ?", email).First(&user)
	return &user
}

func (ur UserRepositoryImpl) FindByUuid(uuidStr string) (u *entity.User, err error) {
	var user entity.User
	if err := infra.Db.Where("uuid = ?", uuidStr).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepositoryImpl) FindByUserName(username string) (u *entity.User, err error)  {
	var user entity.User
	if err := infra.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepositoryImpl) Save(user entity.User) (u *entity.User, err error) {
	if err := infra.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepositoryImpl) Update(user entity.User) (u *entity.User, err error) {
	if err := infra.Db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepositoryImpl) UpdateUserColumn(user entity.User, column string) (u *entity.User, err error) {
	var data string

	switch column {
	case "username":
		data = user.Username
	case "email":
		data = user.Email
	case "password":
		data = user.Password
	}

	if err := infra.Db.Model(&entity.User{}).Where("email = ?", user.Email).Update(column, data).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
