package route

import (
	_controllerAuth "api-station/modules/auth/controller"
	_serviceAuth "api-station/modules/auth/service"
	_repositoryUser "api-station/modules/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(r *gin.Engine, db *gorm.DB) {

	repositoryUser := _repositoryUser.NewrepositoryUser(db)
	serviceAuth := _serviceAuth.NewAuthService(repositoryUser)
	controllerAuth := _controllerAuth.NewAuthController(serviceAuth)
	authRoute := r.Group("/api/v1/auth")
	{
		authRoute.POST("/login", controllerAuth.Login)
		authRoute.POST("/register", controllerAuth.Register)
		authRoute.POST("/email_check", controllerAuth.CheckEmailAvailability)
		authRoute.POST("/username_check", controllerAuth.CheckUsernameAvailability)
	}
}
