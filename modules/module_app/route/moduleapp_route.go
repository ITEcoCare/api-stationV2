package web

import (
	_authMiddleware "api-station/middleware"
	_controllerModuleApp "api-station/modules/module_app/controller"
	_repositoryModuleApp "api-station/modules/module_app/repository"
	_serviceModuleApp "api-station/modules/module_app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ModuleAppRoute(r *gin.Engine, db *gorm.DB) {

	repositoryModuleApp := _repositoryModuleApp.NewModuleAppRepository(db)
	serviceModuleApp := _serviceModuleApp.NewModuleAppService(repositoryModuleApp)
	controllerModuleApp := _controllerModuleApp.NewModuleAppController(serviceModuleApp)

	authMiddleware := _authMiddleware.AuthMiddleware(db)

	moduleAppRoute := r.Group("/api/v1/module-app").Use(authMiddleware)
	{
		moduleAppRoute.POST("/create", controllerModuleApp.Create)
		moduleAppRoute.GET("/read", controllerModuleApp.Read)
		moduleAppRoute.GET("/read/:id", controllerModuleApp.ReadById)
		moduleAppRoute.PUT("/update", controllerModuleApp.Update)
		moduleAppRoute.POST("/delete", controllerModuleApp.Delete)
		moduleAppRoute.GET("/trash", controllerModuleApp.Trash)
		moduleAppRoute.POST("/restore", controllerModuleApp.Restore)
	}
}
