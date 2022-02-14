package web

import (
	_authMiddleware "api-station/middleware"
	_controllerRole "api-station/modules/role/controller"
	_repositoryRole "api-station/modules/role/repository"
	_serviceRole "api-station/modules/role/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRoute(r *gin.Engine, db *gorm.DB) {

	repositoryRole := _repositoryRole.NewRoleRepository(db)
	serviceRole := _serviceRole.NewRoleService(repositoryRole)
	controllerRole := _controllerRole.NewRoleController(serviceRole)

	authMiddleware := _authMiddleware.AuthMiddleware(db)

	roleRoute := r.Group("/api/v1/role").Use(authMiddleware)
	{
		roleRoute.POST("/create", controllerRole.Create)
		roleRoute.GET("/read", controllerRole.Read)
		roleRoute.GET("/read/:id", controllerRole.ReadById)
		roleRoute.PUT("/update", controllerRole.Update)
		roleRoute.POST("/delete", controllerRole.Delete)
		roleRoute.GET("/trash", controllerRole.Trash)
		roleRoute.POST("/restore", controllerRole.Restore)
	}
}
