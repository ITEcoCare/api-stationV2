package repository

import (
	"api-station/models"
	"api-station/modules/user"
	"api-station/response"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewrepositoryUser(db *gorm.DB) user.IRepository {
	return &repositoryUser{db}
}

func (r *repositoryUser) Save(user models.User) response.RepositoryResult {
	log.Print("[repositoryUser]... Save")
	err := r.db.Create(&user).Error

	if err != nil {
		return response.RepositoryResult{
			Result: user,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: user}
}

func (r *repositoryUser) FindAll() response.RepositoryResult {
	log.Print("[repositoryUser]... FindAll")
	var users models.Users
	// Preload("Team").
	err := r.db.Preload("Role").Find(&users).Error
	if err != nil {
		if err != nil {
			return response.RepositoryResult{
				Result: users,
				Error:  err,
			}
		}
	}

	return response.RepositoryResult{Result: users}
}

func (r *repositoryUser) FindById(Id int) response.RepositoryResult {
	log.Print("[repositoryUser]... FindById")
	var user models.User

	err := r.db.Where("id = ?", Id).First(&user).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: user}

}

func (r *repositoryUser) FindByEmail(email string) response.RepositoryResult {
	log.Print("[repositoryUser]... FindByEmail")
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: user}
}

func (r *repositoryUser) FindByUsername(username string) response.RepositoryResult {
	log.Print("[repositoryUser]... FindByUsername")

	var user models.User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: user}

}

func (r *repositoryUser) FindByIdWithRelation(Id int) response.RepositoryResult {
	log.Print("[repositoryUser]... FindById")
	var user models.User
	// Preload("Team").
	err := r.db.Preload("Role").Where("id = ?", Id).First(&user).Error
	if err != nil {
		return response.RepositoryResult{Error: err}
	}

	return response.RepositoryResult{Result: user}

}

func (r *repositoryUser) Update(user models.User) response.RepositoryResult {
	log.Print("[repositoryUser]... Update")

	err := r.db.Save(&user).Error
	if err != nil {
		return response.RepositoryResult{
			Result: user,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: user}
}

func (r *repositoryUser) SoftDelete(user models.User) response.RepositoryResult {
	log.Print("[repositoryUser]... SoftDelete")

	err := r.db.Delete(&user).Error

	if err != nil {
		return response.RepositoryResult{
			Result: user,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: user}
}

func (r *repositoryUser) FindAllWithTrashed() response.RepositoryResult {
	log.Print("[repositoryUser]... FindAllWithTrashed")

	var users models.Users
	// Preload("Team").
	err := r.db.Preload("Role").Unscoped().Where("deleted_at is not null").Find(&users).Error
	if err != nil {
		return response.RepositoryResult{
			Result: users,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: users}
}

func (r *repositoryUser) FindSingleTrashedById(Id int) response.RepositoryResult {
	log.Print("[repositoryUser]... FindAllWithTrashed")

	var user models.User
	err := r.db.Unscoped().Where("id = ?", Id).Where("deleted_at is not null").First(&user).Error
	if err != nil {
		return response.RepositoryResult{
			Result: user,
			Error:  err,
		}
	}

	return response.RepositoryResult{Result: user}
}

func (r repositoryUser) Restore(Id int) response.RepositoryResult {
	log.Print("[repositoryUser]... Restore")
	var user models.User
	err := r.db.Unscoped().Model(&user).Where("id = ?", Id).Update("deleted_at", nil).Error
	fmt.Println(user)
	if err != nil {
		return response.RepositoryResult{
			Result: user,
			Error:  err,
		}
	}
	return response.RepositoryResult{Result: user}
}
