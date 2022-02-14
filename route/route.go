package route

import (
	_authRoute "api-station/auth/route"
	_teamRoute "api-station/team/route"
	_userRoute "api-station/user/route"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB) {

	r := gin.Default()

	r.Static("/images", "./images")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Route Web
	_authRoute.AuthRoute(r, db)
	_userRoute.UserRoute(r, db)
	_teamRoute.TeamRoute(r, db)

	r.Run(":" + os.Getenv("APP_PORT"))

	gin.SetMode(gin.DebugMode)
}