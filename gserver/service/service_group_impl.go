package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var sgsInstance ServiceGroupService

type ServiceGroupServiceImpl struct {
	serviceGroupRepository repository.ServiceGroupRepository
}

func GetServiceGroupServiceInstance() ServiceGroupService {
	if sgsInstance == nil {
		sgsInstance = NewServiceGroupService()
	}
	return sgsInstance
}

func NewServiceGroupService() ServiceGroupService {
	log.Logger.Info("New `ServiceGroupService` instance")
	log.Logger.Info("Inject `ServiceGroupRepository` to `ServiceGroupService`")
	return ServiceGroupServiceImpl{serviceGroupRepository: repository.GetServiceGroupRepositoryInstance(driver.Db)}
}