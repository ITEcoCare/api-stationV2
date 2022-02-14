package web

import (
	_authMiddleware "api-station/middleware"
	_controllerTeam "api-station/modules/team/controller"
	_repositoryTeam "api-station/modules/team/repository"
	_serviceTeam "api-station/modules/team/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TeamRoute(r *gin.Engine, db *gorm.DB) {

	repositoryTeam := _repositoryTeam.NewRepositoryTeam(db)
	serviceTeam := _serviceTeam.NewTeamService(repositoryTeam)
	controllerTeam := _controllerTeam.NewTeamController(serviceTeam)

	authMiddleware := _authMiddleware.AuthMiddleware(db)

	teamRoute := r.Group("/api/v1/team").Use(authMiddleware)
	{
		teamRoute.POST("/create", controllerTeam.Create)
		teamRoute.GET("/read", controllerTeam.Read)
		teamRoute.GET("/read/:id", controllerTeam.ReadById)
		teamRoute.PUT("/update", controllerTeam.Update)
		teamRoute.POST("/delete", controllerTeam.Delete)
		teamRoute.GET("/trash", controllerTeam.Trash)
		teamRoute.POST("/restore", controllerTeam.Restore)
	}
}
