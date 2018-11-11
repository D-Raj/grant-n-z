package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type GroupRepository interface {
	FindAll() (groups []*entity.Group, err error)

	FindByDomain(domain string) *entity.Group

	FindById(id int) (g *entity.Group, err error)

	Save(group entity.Group) (g *entity.Group, err error)
}
