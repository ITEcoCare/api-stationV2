package route

import (
	_controllerAuth "api-station/auth/controller"
	_serviceAuth "api-station/auth/service"
	_repositoryUser "api-station/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(r *gin.Engine, db *gorm.DB) {

	repositoryUser := _repositoryUser.NewrepositoryUser(db)
	serviceAuth := _serviceAuth.NewAuthService(repositoryUser)
	controllerAuth := _controllerAuth.NewAuthController(serviceAuth)
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", controllerAuth.Login)
		auth.POST("/register", controllerAuth.Register)
		auth.POST("/email_check", controllerAuth.CheckEmailAvailability)
		auth.POST("/username_check", controllerAuth.CheckUsernameAvailability)
	}
}
