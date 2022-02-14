package service

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/request"
	"api-station/response"
	"api-station/user"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository user.IRepository
}

func NewAuthService(userRepository user.IRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) CheckLogin(request request.RequestLogin) response.Response {
	log.Print("[authService]... CheckLogin")

	resultUser := s.userRepository.FindByEmail(request.Email)
	if resultUser.Error != nil {
		return response.Response{Success: false, Message: resultUser.Error.Error()}
	}

	user := resultUser.Result.(models.User)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return response.Response{Success: false, Message: err.Error()}
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return response.Response{Success: false, Message: err.Error()}
	}

	responseAuthUser := response.ResponseUserAuth{
		ID:              user.ID,
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Token:           token,
	}

	return response.Response{Success: true, Message: "User Detail", Data: responseAuthUser}
}

func (s *authService) Register(request models.User) response.Response {
	log.Print("[authService]... Register")

	t := time.Now().UTC()
	timeNow := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	user := models.User{}
	user.RoleId = request.RoleId
	user.TeamId = request.TeamId
	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	user.EmailVerifiedAt = &timeNow

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Response{
			Success: false,
			Message: err.Error(),
		}
	}

	user.Password = string(passwordHash)

	resultUser := s.userRepository.Save(user)

	if resultUser.Error != nil {
		return response.Response{Success: false, Message: resultUser.Error.Error()}
	}

	newUser := resultUser.Result.(models.User)
	userDetail := s.userRepository.FindByIdWithRelation(newUser.ID)

	// map interface to models
	resData := userDetail.Result.(models.User)
	return response.Response{Success: true, Message: "User has been registered", Data: resData}

}

func (s *authService) CheckEmailAvailability(email string) response.Response {

	log.Print("[authService]... CheckEmailAvailability")

	resultUser := s.userRepository.FindByEmail(email)

	message := "Email is available"
	if resultUser.Result != nil {
		message = "Email has been registered"
	}

	return response.Response{Success: true, Message: message, Data: nil}
}

func (s *authService) CheckUsernameAvailability(username string) response.Response {

	log.Print("[authService]... CheckUsernameAvailability")

	resultUser := s.userRepository.FindByUsername(username)

	message := "Username is available"
	if resultUser.Result != nil {
		message = "Username has been registered"
	}

	return response.Response{Success: true, Message: message, Data: nil}
}
