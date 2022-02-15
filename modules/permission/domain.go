package permission

import (
	"api-station/models"
	"api-station/response"
)

type IRepository interface {
	Save(permission models.Permission) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	Update(permission models.Permission) response.RepositoryResult
	SoftDelete(permission models.Permission) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(permission models.Permission) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	Update(permission models.Permission) response.Response
	Delete(Id int) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
