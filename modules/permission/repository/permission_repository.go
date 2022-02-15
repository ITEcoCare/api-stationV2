package repository

import (
	"api-station/models"
	"api-station/modules/permission"
	"api-station/response"
	"log"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) permission.IRepository {
	return &permissionRepository{db}
}

func (r *permissionRepository) Save(permission models.Permission) response.RepositoryResult {
	log.Print("[permissionRepository]... Save")
	err := r.db.Create(&permission).Error

	if err != nil {
		return response.RepositoryResult{
			Result: permission,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permission}
}

func (r *permissionRepository) FindAll() response.RepositoryResult {
	log.Print("[permissionRepository]... FindAll")
	var permissions models.Permissions

	err := r.db.Find(&permissions).Error
	if err != nil {
		return response.RepositoryResult{
			Result: permissions,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permissions}
}

func (r *permissionRepository) FindById(Id int) response.RepositoryResult {
	log.Print("[permissionRepository]... FindById")
	var permission models.Permission

	err := r.db.Where("id = ?", Id).First(&permission).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: permission}

}

func (r *permissionRepository) Update(permission models.Permission) response.RepositoryResult {
	log.Print("[permissionRepository]... Update")

	err := r.db.Save(&permission).Error
	if err != nil {
		return response.RepositoryResult{
			Result: permission,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permission}
}

func (r *permissionRepository) SoftDelete(permission models.Permission) response.RepositoryResult {
	log.Print("[permissionRepository]... SoftDelete")

	err := r.db.Delete(&permission).Error

	if err != nil {
		return response.RepositoryResult{
			Result: permission,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permission}
}

func (r *permissionRepository) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[permissionRepository]... FindAllWithTrashed")

	var permissions models.Permissions

	err := r.db.Unscoped().Where("deleted_at is not null").Find(&permissions).Error
	if err != nil {
		return response.RepositoryResult{
			Result: permissions,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permissions}
}

func (r *permissionRepository) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[permissionRepository]... FindAllWithTrashed")

	var permission models.Permission
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&permission).Error
	if err != nil {
		return response.RepositoryResult{
			Result: permission,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: permission}
}

func (r permissionRepository) Restore(Id int) response.RepositoryResult {
	log.Print("[permissionRepository]... Restore")
	var permission models.Permission
	err := r.db.Unscoped().Model(&permission).Where("id = ?", Id).Update("deleted_at", nil).Error

	if err != nil {
		return response.RepositoryResult{
			Result: permission,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: permission}
}
