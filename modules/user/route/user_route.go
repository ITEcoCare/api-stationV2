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

	userRoute := r.Group("/api/v1/user").Use(authMiddleware)
	{
		userRoute.POST("/create", controllerUser.Create)
		userRoute.GET("/read", controllerUser.Read)
		userRoute.GET("/read/:id", controllerUser.ReadById)
		userRoute.PUT("/update", controllerUser.Update)
		userRoute.POST("/delete", controllerUser.Delete)
		userRoute.GET("/trash", controllerUser.Trash)
		userRoute.POST("/restore", controllerUser.Restore)
		// userAdmin.POST("/upload_avatar", handler.UploadAvatar)
		// userAdmin.POST("/restore", handler.RestoreUser)
	}
}
