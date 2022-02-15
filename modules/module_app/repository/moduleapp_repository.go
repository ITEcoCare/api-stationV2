package repository

import (
	"api-station/models"
	"api-station/modules/module_app"
	"api-station/response"
	"log"

	"gorm.io/gorm"
)

type moduleAppRepository struct {
	db *gorm.DB
}

func NewModuleAppRepository(db *gorm.DB) module_app.IRepository {
	return &moduleAppRepository{db}
}

func (r *moduleAppRepository) Save(module models.ModuleApp) response.RepositoryResult {
	log.Print("[moduleAppRepository]... Save")
	err := r.db.Create(&module).Error

	if err != nil {
		return response.RepositoryResult{
			Result: module,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: module}
}

func (r *moduleAppRepository) FindAll() response.RepositoryResult {
	log.Print("[moduleAppRepository]... FindAll")
	var modules models.ModuleApps

	err := r.db.Find(&modules).Error
	if err != nil {
		return response.RepositoryResult{
			Result: modules,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: modules}
}

func (r *moduleAppRepository) FindById(Id int) response.RepositoryResult {
	log.Print("[moduleAppRepository]... FindById")
	var module models.ModuleApp

	err := r.db.Where("id = ?", Id).First(&module).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: module}
}

func (r *moduleAppRepository) Update(module models.ModuleApp) response.RepositoryResult {
	log.Print("[moduleAppRepository]... Update")

	err := r.db.Save(&module).Error
	if err != nil {
		return response.RepositoryResult{
			Result: module,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: module}
}

func (r *moduleAppRepository) SoftDelete(module models.ModuleApp) response.RepositoryResult {
	log.Print("[moduleAppRepository]... SoftDelete")

	err := r.db.Delete(&module).Error

	if err != nil {
		return response.RepositoryResult{
			Result: module,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: module}
}

func (r *moduleAppRepository) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[moduleAppRepository]... FindAllWithTrashed")

	var modules models.ModuleApps

	err := r.db.Unscoped().Where("deleted_at is not null").Find(&modules).Error
	if err != nil {
		return response.RepositoryResult{
			Result: modules,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: modules}
}

func (r *moduleAppRepository) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[moduleAppRepository]... FindSingleTrashedById")

	var module models.ModuleApp
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&module).Error
	if err != nil {
		return response.RepositoryResult{
			Result: module,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: module}
}

func (r moduleAppRepository) Restore(Id int) response.RepositoryResult {
	log.Print("[moduleAppRepository]... Restore")
	var module models.ModuleApp
	err := r.db.Unscoped().Model(&module).Where("id = ?", Id).Update("deleted_at", nil).Error

	if err != nil {
		return response.RepositoryResult{
			Result: module,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: module}
}
