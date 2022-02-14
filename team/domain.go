package team

import (
	"api-station/models"
	"api-station/request"
	"api-station/response"
)

type IRepository interface {
	Save(team models.Team) response.RepositoryResult
	FindAll() response.RepositoryResult
	FindById(Id int) response.RepositoryResult
	Update(team models.Team) response.RepositoryResult
	SoftDelete(team models.Team) response.RepositoryResult
	FindAllWithTrashed() response.RepositoryResult
	FindSingleTrashedById(Id int) response.RepositoryResult
	Restore(Id int) response.RepositoryResult
}

type IService interface {
	Create(team models.Team) response.Response
	Read() response.Response
	ReadById(Id int) response.Response
	Update(team models.Team) response.Response
	Delete(requestId request.RequestId) response.Response
	Trash() response.Response
	Restore(Id int) response.Response
}
