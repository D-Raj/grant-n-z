package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type GroupRepository interface {
	FindAll() (groups []*entity.Group, err error)

	FindByDomain(domain string) *entity.Group

	Save(group entity.Group) (g *entity.Group, err error)
}
