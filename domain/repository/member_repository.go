package repository

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type MemberRepository interface {
	FindByUserUuidAndServiceUuid(userUuid uuid.UUID, serviceUuid uuid.UUID) *entity.Member

	FindByUserUuid(userUuid uuid.UUID) []*entity.Member

	Save(member entity.Member) *entity.Member
}
