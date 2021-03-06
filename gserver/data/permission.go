package data

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var prInstance PermissionRepository

type PermissionRepository interface {
	// Find all permission
	FindAll() ([]*entity.Permission, *model.ErrorResBody)

	// Find permission by id
	FindById(id int) (*entity.Permission, *model.ErrorResBody)

	// Find permission by name
	FindByName(name string) (*entity.Permission, *model.ErrorResBody)

	// Find permission by name array
	FindByNames(names []string) ([]*entity.Permission, *model.ErrorResBody)

	// Find permissions by group id
	// Join group_permission and permission
	FindByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody)

	// Find permission name by id
	FindNameById(id int) *string

	// Save permission
	Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody)

	// Save permission with relational data
	SaveWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody)
}

type PermissionRepositoryImpl struct {
	Db *gorm.DB
}

func GetPermissionRepositoryInstance(db *gorm.DB) PermissionRepository {
	if prInstance == nil {
		prInstance = NewPermissionRepository(db)
	}
	return prInstance
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	log.Logger.Info("New `PermissionRepository` instance")
	return PermissionRepositoryImpl{Db: db}
}

func (pri PermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	if err := pri.Db.Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	if err := pri.Db.Where("id = ?", id).Find(&permission).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) FindByName(name string) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	if err := pri.Db.Where("name = ?", name).Find(&permission).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) FindByNames(names []string) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	if err := pri.Db.Where("name IN (?)", names).Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission

	if err := pri.Db.Table(entity.GroupPermissionTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PermissionTable.String(),
			entity.GroupPermissionTable.String(),
			entity.GroupPermissionPermissionId,
			entity.PermissionTable.String(),
			entity.PermissionId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.GroupPermissionTable.String(),
			entity.GroupPermissionGroupId), groupId).
		Scan(&permissions).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found permission")
		}

		return nil, model.InternalServerError()
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindNameById(id int) *string {
	if id == 0 {
		return nil
	}
	permission, err := pri.FindById(id)
	if err != nil {
		return nil
	}
	return &permission.Name
}

func (pri PermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	if err := pri.Db.Create(&permission).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) SaveWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	tx := pri.Db.Begin()

	// Save permission
	if err := tx.Create(&permission).Error; err != nil {
		log.Logger.Warn("Failed to save permissions at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit permissions data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_permissions
	groupPermission := entity.GroupPermission{
		PermissionId:  permission.Id,
		GroupId: groupId,
	}
	if err := tx.Create(&groupPermission).Error; err != nil {
		log.Logger.Warn("Failed to save group_permissions at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_permissions data.")
		}

		return nil, model.InternalServerError()
	}

	tx.Commit()

	return &permission, nil
}
