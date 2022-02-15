package web

import (
	_authMiddleware "api-station/middleware"
	_controllerPermission "api-station/modules/permission/controller"
	_repositoryPermission "api-station/modules/permission/repository"
	_servicePermission "api-station/modules/permission/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionRoute(r *gin.Engine, db *gorm.DB) {

	repositoryPermission := _repositoryPermission.NewPermissionRepository(db)
	servicePermission := _servicePermission.NewPermissionService(repositoryPermission)
	controllerPermission := _controllerPermission.NewPermissionController(servicePermission)

	authMiddleware := _authMiddleware.AuthMiddleware(db)

	permissionRoute := r.Group("/api/v1/permission").Use(authMiddleware)
	{
		permissionRoute.POST("/create", controllerPermission.Create)
		permissionRoute.GET("/read", controllerPermission.Read)
		permissionRoute.GET("/read/:id", controllerPermission.ReadById)
		permissionRoute.PUT("/update", controllerPermission.Update)
		permissionRoute.POST("/delete", controllerPermission.Delete)
		permissionRoute.GET("/trash", controllerPermission.Trash)
		permissionRoute.POST("/restore", controllerPermission.Restore)
	}
}
