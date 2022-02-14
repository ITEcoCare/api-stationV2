package repository

import (
	"api-station/models"
	"api-station/modules/team"
	"api-station/response"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type repositoryTeam struct {
	db *gorm.DB
}

func NewRepositoryTeam(db *gorm.DB) team.IRepository {
	return &repositoryTeam{db}
}

func (r *repositoryTeam) Save(team models.Team) response.RepositoryResult {
	log.Print("[repositoryTeam]... Save")
	err := r.db.Create(&team).Error

	if err != nil {
		return response.RepositoryResult{
			Result: team,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: team}
}

func (r *repositoryTeam) FindAll() response.RepositoryResult {
	log.Print("[repositoryTeam]... FindAll")
	var teams models.Teams

	err := r.db.Find(&teams).Error
	if err != nil {
		return response.RepositoryResult{
			Result: teams,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: teams}
}

func (r *repositoryTeam) FindById(Id int) response.RepositoryResult {
	log.Print("[repositoryTeam]... FindById")
	var team models.Team

	err := r.db.Where("id = ?", Id).First(&team).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: team}

}

func (r *repositoryTeam) Update(team models.Team) response.RepositoryResult {
	log.Print("[repositoryTeam]... Update")

	err := r.db.Save(&team).Error
	if err != nil {
		return response.RepositoryResult{
			Result: team,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: team}
}

func (r *repositoryTeam) SoftDelete(team models.Team) response.RepositoryResult {
	log.Print("[repositoryTeam]... SoftDelete")

	err := r.db.Delete(&team).Error

	if err != nil {
		return response.RepositoryResult{
			Result: team,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: team}
}

func (r *repositoryTeam) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[repositoryTeam]... FindAllWithTrashed")

	var teams models.Teams

	err := r.db.Unscoped().Where("deleted_at is not null").Find(&teams).Error
	if err != nil {
		return response.RepositoryResult{
			Result: teams,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: teams}
}

func (r *repositoryTeam) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[repositoryTeam]... FindAllWithTrashed")

	var team models.Team
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&team).Error
	if err != nil {
		return response.RepositoryResult{
			Result: team,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: team}
}

func (r repositoryTeam) Restore(Id int) response.RepositoryResult {
	log.Print("[repositoryTeam]... Restore")
	var team models.Team
	err := r.db.Unscoped().Model(&team).Where("id = ?", Id).Update("deleted_at", nil).Error
	fmt.Println(team)
	if err != nil {
		return response.RepositoryResult{
			Result: team,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: team}
}
