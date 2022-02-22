package company

import (
	"api-station/models"
	"api-station/response"
)

type IRepository interface {
	Save(company models.Company) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	Update(company models.Company) response.RepositoryResult
	SoftDelete(company models.Company) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(comapny models.Company) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	Update(user models.Company) response.Response
	Delete(Id int) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
