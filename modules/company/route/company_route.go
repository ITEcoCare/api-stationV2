package route

import (
	_controllerCompany "api-station/modules/company/controller"
	_repositoryCompany "api-station/modules/company/repository"
	_serviceCompany "api-station/modules/company/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CompanyRoute(r *gin.Engine, db *gorm.DB) {

	repo := _repositoryCompany.NewCompanyRepository(db)
	service := _serviceCompany.NewServiceCompany(repo)
	controller := _controllerCompany.NewCompanyController(service)

	// authMiddleware := _authMiddleware.AuthMiddleware(db)
	// .Use(authMiddleware)
	userRoute := r.Group("/api/v1/company")
	{
		userRoute.POST("/create", controller.Create)
		userRoute.GET("/read", controller.Read)
		userRoute.GET("/read/:id", controller.ReadById)
		userRoute.PUT("/update", controller.Update)
		userRoute.POST("/delete", controller.Delete)
		userRoute.GET("/trash", controller.Trash)
		userRoute.POST("/restore", controller.Restore)
		// userAdmin.POST("/upload_avatar", handler.UploadAvatar)
		// userAdmin.POST("/restore", handler.RestoreUser)
	}
}
