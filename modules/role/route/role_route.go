package web

import (
	_Middleware "api-station/middleware"
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

	authMiddleware := _Middleware.AuthMiddleware(db)

	roleRoute := r.Group("/api/v1/role").Use(authMiddleware)
	{
		roleRoute.POST("/create", _Middleware.PermissionMiddleware("create-role", db), controllerRole.Create)
		roleRoute.GET("/read", _Middleware.PermissionMiddleware("read-role", db), controllerRole.Read)
		roleRoute.GET("/read/:id", _Middleware.PermissionMiddleware("read-role", db), controllerRole.ReadById)
		roleRoute.PUT("/update", _Middleware.PermissionMiddleware("update-role", db), controllerRole.Update)
		roleRoute.POST("/delete", _Middleware.PermissionMiddleware("delete-role", db), controllerRole.Delete)
		roleRoute.GET("/trash", _Middleware.PermissionMiddleware("read-role", db), controllerRole.Trash)
		roleRoute.POST("/restore", _Middleware.PermissionMiddleware("update-role", db), controllerRole.Restore)
	}
}
