package module_app

import (
	"api-station/models"
	"api-station/response"
)

type IRepository interface {
	Save(module models.ModuleApp) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	Update(module models.ModuleApp) response.RepositoryResult
	SoftDelete(module models.ModuleApp) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(module models.ModuleApp) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	Update(module models.ModuleApp) response.Response
	Delete(Id int) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
