package middleware

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/response"
	"fmt"

	// "api-station/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// userService user.IService
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			res := response.Response{Success: false, Message: "Unauthorized"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		tokenString := ""
		arrTokenSplt := strings.Split(authHeader, " ")
		if len(arrTokenSplt) == 2 {
			tokenString = arrTokenSplt[1]
		}

		token, err := helpers.ValidateToken(tokenString)

		if err != nil {
			res := response.Response{Success: false, Message: "Unauthorized, Invalid Token"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			res := response.Response{Success: false, Message: "Unauthorized, Invalid Token2"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userId := int(claim["user_id"].(float64))

		var user models.User
		// user := userService.ReadById(userId)
		// Preload("Team").
		result := db.Preload("Role").Where("id = ?", userId).First(&user).Error
		if result != nil {
			res := response.Response{Success: false, Message: "Unauthorized"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		c.Set("currentUser", user)
	}
}

func PermissionMiddleware(permissionName string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(models.User)
		fmt.Println("PermissionMiddleware")
		fmt.Println(currentUser)
		fmt.Println(currentUser.Role)
		// validation user permissions exist
		// fmt.Println(currentUser.Team)
	}
}
