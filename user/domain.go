package user

import (
	"api-station/models"
	"api-station/request"
	"api-station/response"
)

type IRepository interface {
	Save(user models.User) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	FindByEmail(email string) response.RepositoryResult
	FindByUsername(username string) response.RepositoryResult
	FindByIdWithRelation(Id int) response.RepositoryResult
	Update(user models.User) response.RepositoryResult
	SoftDelete(user models.User) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(user models.User) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	ReadByEmail(email string) response.Response
	ReadByUsername(username string) response.Response
	ReadByIdWithRelation(Id int) response.Response
	Update(user models.User) response.Response
	Delete(requestId request.RequestId) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
