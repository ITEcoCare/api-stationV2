package web

import (
	_Middleware "api-station/middleware"
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

	authMiddleware := _Middleware.AuthMiddleware(db)

	teamRoute := r.Group("/api/v1/team").Use(authMiddleware)
	{
		teamRoute.POST("/create", _Middleware.PermissionMiddleware("create-team", db), controllerTeam.Create)
		teamRoute.GET("/read", _Middleware.PermissionMiddleware("read-team", db), controllerTeam.Read)
		teamRoute.GET("/read/:id", _Middleware.PermissionMiddleware("create-team", db), controllerTeam.ReadById)
		teamRoute.PUT("/update", _Middleware.PermissionMiddleware("update-team", db), controllerTeam.Update)
		teamRoute.POST("/delete", _Middleware.PermissionMiddleware("delete-team", db), controllerTeam.Delete)
		teamRoute.GET("/trash", _Middleware.PermissionMiddleware("read-team", db), controllerTeam.Trash)
		teamRoute.POST("/restore", _Middleware.PermissionMiddleware("update-team", db), controllerTeam.Restore)
	}
}
