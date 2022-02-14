package repository

import (
	"api-station/models"
	"api-station/modules/role"
	"api-station/response"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.IRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Save(role models.Role) response.RepositoryResult {
	log.Print("[roleRepository]... Save")
	err := r.db.Create(&role).Error

	if err != nil {
		return response.RepositoryResult{
			Result: role,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: role}
}

func (r *roleRepository) FindAll() response.RepositoryResult {
	log.Print("[roleRepository]... FindAll")
	var roles models.Roles

	err := r.db.Find(&roles).Error
	if err != nil {
		return response.RepositoryResult{
			Result: roles,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: roles}
}

func (r *roleRepository) FindById(Id int) response.RepositoryResult {
	log.Print("[roleRepository]... FindById")
	var role models.Role

	err := r.db.Where("id = ?", Id).First(&role).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: role}

}

func (r *roleRepository) Update(role models.Role) response.RepositoryResult {
	log.Print("[roleRepository]... Update")

	err := r.db.Save(&role).Error
	if err != nil {
		return response.RepositoryResult{
			Result: role,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: role}
}

func (r *roleRepository) SoftDelete(role models.Role) response.RepositoryResult {
	log.Print("[roleRepository]... SoftDelete")

	err := r.db.Delete(&role).Error

	if err != nil {
		return response.RepositoryResult{
			Result: role,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: role}
}

func (r *roleRepository) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[roleRepository]... FindAllWithTrashed")

	var roles models.Roles

	err := r.db.Unscoped().Where("deleted_at is not null").Find(&roles).Error
	if err != nil {
		return response.RepositoryResult{
			Result: roles,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: roles}
}

func (r *roleRepository) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[roleRepository]... FindAllWithTrashed")

	var role models.Role
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&role).Error
	if err != nil {
		return response.RepositoryResult{
			Result: role,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: role}
}

func (r roleRepository) Restore(Id int) response.RepositoryResult {
	log.Print("[roleRepository]... Restore")
	var role models.Role
	err := r.db.Unscoped().Model(&role).Where("id = ?", Id).Update("deleted_at", nil).Error
	fmt.Println(role)
	if err != nil {
		return response.RepositoryResult{
			Result: role,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: role}
}
