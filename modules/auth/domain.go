package auth

import (
	"api-station/models"
	"api-station/request"
	"api-station/response"
)

type IService interface {
	Register(user models.User) response.Response
	CheckLogin(request request.RequestLogin) response.Response
	CheckEmailAvailability(email string) response.Response
	CheckUsernameAvailability(username string) response.Response
}
