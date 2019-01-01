package service

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/infra"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	userRepository repository.UserRepository
	principalRepository repository.PrincipalRepository
}

func (us UserService) NewUserService() UserService {
	us.userRepository = repository.UserRepositoryImpl{}
	return us
}

func (us UserService) GenerateUser(request *entity.UserReq) (response *entity.User, err error, code int) {
	infra.Db = infra.Db.Begin()

	user := entity.User {
		Uuid: uuid.NewV4(),
		Email: request.Email,
		Password: us.EncryptPw(request.Password),
	}

	if us.userRepository.FindByEmail(user.Email).Email != "" {
		err = errors.New("already exist email.")
		return nil, err, http.StatusConflict
	}

	response, err = us.userRepository.Save(user)
	if err != nil {
		err = errors.New("user insert error.")
		return nil, err, http.StatusInternalServerError
	}

	principal := entity.Principal{
		UserId: response.Id,
		GroupId:
	}

	return
}

func (us UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}
//
//func (u UserService) GetUserByEmail(email string) *entity.User {
//	return u.UserRepository.FindByEmail(email)
//}
//
//func (u UserService) GetUserByUuid(uuid string) *entity.User {
//	return u.UserRepository.FindByUuid(uuid)
//}
//
//func (u UserService) InsertUser(user entity.User) *entity.User {
//	user.Uuid, _ = uuid.NewV4()
//	user.Password = u.EncryptPw(user.Password)
//	return u.UserRepository.Save(user)
//}
//
//func (u UserService) UpdateUser(user entity.User) *entity.User {
//	user.Password = u.EncryptPw(user.Password)
//	return u.UserRepository.Update(user)
//}
//
//func (u UserService) UpdateUserColumn(user entity.User, column string) *entity.User {
//	user.Password = u.EncryptPw(user.Password)
//	return u.UserRepository.UpdateUserColumn(user, column)
//}
//
//func (u UserService) PostUserData(user *entity.User) *handler.ErrorResponse {
//	userData := u.GetUserByEmail(user.Email)
//	if userData == nil {
//		return handler.InternalServerError("")
//	}
//
//	if len(userData.Email) > 0 {
//		return handler.Conflict("")
//	}
//
//	userData = u.InsertUser(*user)
//	if userData == nil {
//		return handler.InternalServerError("")
//	}
//
//	return nil
//}
//
//func (u UserService) PutUserColumnData(user *entity.User, column string) *handler.ErrorResponse {
//	userData := u.GetUserByEmail(user.Email)
//	if userData == nil {
//		return handler.InternalServerError("")
//	}
//
//	if len(userData.Email) == 0 {
//		return handler.NotFound("")
//	}
//
//	userData = u.UpdateUserColumn(*user, column)
//	if userData == nil {
//		return handler.InternalServerError("")
//	}
//
//	return nil
//}