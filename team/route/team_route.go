package web

import (
	_controllerTeam "api-station/team/controller"
	_repositoryTeam "api-station/team/repository"
	_serviceTeam "api-station/team/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TeamRoute(r *gin.Engine, db *gorm.DB) {

	repositoryTeam := _repositoryTeam.NewRepositoryTeam(db)
	serviceTeam := _serviceTeam.NewTeamService(repositoryTeam)
	controllerTeam := _controllerTeam.NewTeamController(serviceTeam)

	// authMiddleware := _authMiddleware.AuthMiddleware(serviceTeam)

	userTeam := r.Group("/api/v1/team")
	{
		userTeam.POST("/create", controllerTeam.Create)
		userTeam.GET("/read", controllerTeam.Read)
		userTeam.GET("/read/:id", controllerTeam.ReadById)
		userTeam.PUT("/update", controllerTeam.Update)
		userTeam.POST("/delete", controllerTeam.Delete)
		userTeam.GET("/trash", controllerTeam.Trash)
		userTeam.POST("/restore", controllerTeam.Restore)
	}
}
