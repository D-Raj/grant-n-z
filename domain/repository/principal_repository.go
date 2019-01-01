package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type PrincipalRepository interface {
	Save(principal entity.Principal) (p *entity.Principal, err error)
}
