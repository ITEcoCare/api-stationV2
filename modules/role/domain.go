package role

import (
	"api-station/models"
	"api-station/request"
	"api-station/response"
)

type IRepository interface {
	Save(role models.Role) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	Update(role models.Role) response.RepositoryResult
	SoftDelete(role models.Role) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(role models.Role) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	Update(role models.Role) response.Response
	Delete(requestId request.RequestId) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
