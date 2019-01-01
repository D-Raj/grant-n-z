package service
//
//import (
//	"github.com/satori/go.uuid"
//	"github.com/tomoyane/grant-n-z/domain/entity"
//	"github.com/tomoyane/grant-n-z/domain/repository"
//	"github.com/tomoyane/grant-n-z/handler"
//)
//
//type RoleService struct {
//	RoleRepository repository.RoleRepository
//}
//
//func (r RoleService) GetRoleByUserUuid(userUuid string) *entity.Role {
//	return r.RoleRepository.FindByUserUuid(userUuid)
//}
//
//func (r RoleService) GetRoleByPermission(permission string) *entity.Role {
//	return r.RoleRepository.FindByPermission(permission)
//}
//
//func (r RoleService) InsertRole(role entity.Role) *entity.Role {
//	role.Uuid, _ = uuid.NewV4()
//	return r.RoleRepository.Save(role)
//}
//
//func (r RoleService) PostRoleData(role *entity.Role) (insertedRole *entity.Role, errRes *handler.ErrorResponse) {
//	roleData := r.GetRoleByPermission(role.Permission)
//	if roleData == nil {
//		return nil, handler.InternalServerError("aa")
//	}
//
//	if len(roleData.Permission) > 0 {
//		return nil, handler.Conflict("")
//	}
//
//	roleData = r.InsertRole(*role)
//	if roleData == nil {
//		return nil, handler.InternalServerError("aa")
//	}
//
//	return roleData, nil
//}