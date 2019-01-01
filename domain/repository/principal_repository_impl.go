package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type PrincipalRepositoryImpl struct {
}

func (pr PrincipalRepositoryImpl) NewPrincipalRepository() PrincipalRepository {
	return pr
}

func (pr PrincipalRepositoryImpl) Save(principal entity.Principal) (p *entity.Principal, err error) {
	if err := infra.Db.Create(&principal).Error; err != nil {
		return nil, err
	}
	return &principal, nil
}
