package middleware

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/response"

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
			res := response.Response{Success: false, Message: "Unauthorized1"}
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
			res := response.Response{Success: false, Message: "Invalid Token"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			res := response.Response{Success: false, Message: "Unauthorized3"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userId := int(claim["user_id"].(float64))
		var user models.User
		// user := userService.ReadById(userId)
		result := db.Where("id = ?", userId).First(&user).Error
		if result != nil {
			res := response.Response{Success: false, Message: "Unauthorized"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		c.Set("currentUser", user)
	}
}
