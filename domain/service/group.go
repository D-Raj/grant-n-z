package service

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

type GroupService struct {
	groupRepository repository.GroupRepository
}

func (gs GroupService) NewGroupService() GroupService {
	gs.groupRepository = repository.GroupRepositoryImpl{}
	return gs
}

func (gs GroupService) GetAll() (response []*entity.Group, err error) {
	groups, err := gs.groupRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (gs GroupService) Insert(request *entity.Group) (response *entity.Group, err error, code int) {
	group := gs.groupRepository.FindByDomain(request.Domain)
	if len(group.Domain) != 0 {
		err = errors.New("already exist resource.")
		return nil, err, http.StatusConflict
	}

	group, err = gs.groupRepository.Save(*request)
	if err != nil {
		err = errors.New("insert error.")
		return nil, err, http.StatusInternalServerError
	}
	return group, nil, http.StatusCreated
}
