package web

import (
	_authMiddleware "api-station/middleware"
	_controllerUser "api-station/modules/user/controller"
	_repositoryUser "api-station/modules/user/repository"
	_serviceUser "api-station/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {

	repositoryUser := _repositoryUser.NewrepositoryUser(db)
	serviceUser := _serviceUser.NewUserService(repositoryUser)
	controllerUser := _controllerUser.NewUserController(serviceUser)

	authMiddleware := _authMiddleware.AuthMiddleware(db)

	userAdmin := r.Group("/api/v1/user").Use(authMiddleware)
	{
		userAdmin.POST("/create", controllerUser.Create)
		userAdmin.GET("/read", controllerUser.Read)
		userAdmin.GET("/read/:id", controllerUser.ReadById)
		userAdmin.PUT("/update", controllerUser.Update)
		userAdmin.POST("/delete", controllerUser.Delete)
		userAdmin.GET("/trash", controllerUser.Trash)
		userAdmin.POST("/restore", controllerUser.Restore)
		// userAdmin.POST("/upload_avatar", handler.UploadAvatar)
		// userAdmin.POST("/restore", handler.RestoreUser)
	}
}
