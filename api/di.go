package api

import "github.com/tomoyane/grant-n-z/domain/service"

func NewGroupService() service.GroupService {
	return service.GroupService{}.NewGroupService()
}

func NewUserService() service.UserService {
	return service.UserService{}.NewUserService()
}
