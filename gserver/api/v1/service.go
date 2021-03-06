package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var sInstance Service

type Service interface {
	// Http GET method
	// All service data
	// Endpoint is `/api/v1/services`
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	// Add user to not main service
	// Endpoint is `/api/v1/services/add_user`
	Post(w http.ResponseWriter, r *http.Request)
}

type ServiceImpl struct {
	ServiceService     service.Service
	UserServiceService service.UserServiceService
	TokenService       service.TokenService
}

func GetServiceInstance() Service {
	if sInstance == nil {
		sInstance = NewService()
	}
	return sInstance
}

func NewService() Service {
	log.Logger.Info("New `v1.Service` instance")
	return ServiceImpl{
		ServiceService:     service.GetServiceInstance(),
		UserServiceService: service.GetUserServiceServiceInstance(),
		TokenService:       service.GetTokenServiceInstance(),
	}
}

func (s ServiceImpl) Get(w http.ResponseWriter, r *http.Request) {
	services, err := s.ServiceService.GetServices()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
	}

	res, _ := json.Marshal(services)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s ServiceImpl) Post(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}
	userEntity.Username = userEntity.Email
	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	// Authentication user
	token, err := s.TokenService.Generate("", "", *userEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}
	token = "Bearer " + token
	authUser, err := s.TokenService.GetAuthUserInToken(token)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	// Check exist service
	serviceEntity, err := s.ServiceService.GetServiceOfApiKey()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	// Insert user_services
	userServiceEntity := &entity.UserService{
		UserId:    authUser.UserId,
		ServiceId: serviceEntity.Id,
	}
	userService, err := s.UserServiceService.InsertUserService(*userServiceEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userService)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
