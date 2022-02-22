package repository

import (
	"api-station/models"
	"api-station/modules/company"
	"api-station/response"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) company.IRepository {
	return &companyRepository{db}
}

func (r *companyRepository) Save(company models.Company) response.RepositoryResult {
	log.Print("[companyRepository]... Save")
	err := r.db.Create(&company).Error

	if err != nil {
		return response.RepositoryResult{
			Result: company,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: company}
}

func (r *companyRepository) FindAll() response.RepositoryResult {
	log.Print("[companyRepository]... FindAll")
	var companies models.Companies

	err := r.db.Find(&companies).Error
	if err != nil {
		if err != nil {
			return response.RepositoryResult{
				Result: companies,
				Error:  err,
			}
		}
	}

	return response.RepositoryResult{Result: companies}
}

func (r *companyRepository) FindById(Id int) response.RepositoryResult {
	log.Print("[companyRepository]... FindById")
	var company models.Company

	err := r.db.Where("id = ?", Id).First(&company).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: company}

}

func (r *companyRepository) Update(company models.Company) response.RepositoryResult {
	log.Print("[companyRepository]... Update")

	err := r.db.Save(&company).Error
	if err != nil {
		return response.RepositoryResult{
			Result: company,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: company}
}

func (r *companyRepository) SoftDelete(company models.Company) response.RepositoryResult {
	log.Print("[companyRepository]... SoftDelete")

	err := r.db.Delete(&company).Error

	if err != nil {
		return response.RepositoryResult{
			Result: company,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: company}
}

func (r *companyRepository) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[companyRepository]... FindAllWithTrashed")

	var companies models.Companies

	err := r.db.Unscoped().Where("deleted_at is not null").Find(&companies).Error
	if err != nil {
		return response.RepositoryResult{
			Result: companies,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: companies}
}

func (r *companyRepository) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[companyRepository]... FindSingleTrashedById")

	var company models.Company
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&company).Error
	if err != nil {
		return response.RepositoryResult{
			Result: company,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: company}
}

func (r companyRepository) Restore(Id int) response.RepositoryResult {
	log.Print("[companyRepository]... Restore")
	var company models.Company
	err := r.db.Unscoped().Model(&company).Where("id = ?", Id).Update("deleted_at", nil).Error
	fmt.Println(company)
	if err != nil {
		return response.RepositoryResult{
			Result: company,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: company}
}
