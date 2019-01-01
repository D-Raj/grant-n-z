package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type GroupRepositoryImpl struct {
}

func (gr GroupRepositoryImpl) NewGroupRepository() GroupRepository {
	return gr
}

func (gr GroupRepositoryImpl) FindByDomain(domain string) *entity.Group {
	var group entity.Group
	infra.Db.Where("domain = ?", domain).First(&group)
	return &group
}

func (gr GroupRepositoryImpl) FindById(id int) (g *entity.Group, err error) {
	var group entity.Group
	if err := infra.Db.Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (gr GroupRepositoryImpl) FindAll() (groups []*entity.Group, err error) {
	if err := infra.Db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (gr GroupRepositoryImpl) Save(group entity.Group) (response *entity.Group, err error) {
	if err := infra.Db.Create(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}
