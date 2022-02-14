package service

import (
	"api-station/models"
	"api-station/modules/user"
	"api-station/request"
	"api-station/response"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.IRepository
}

func NewUserService(userRepository user.IRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) Create(request models.User) response.Response {
	log.Print("[userService]...Create")

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

func (s *userService) Read() response.Response {
	log.Print("[userService]...Read")

	resultUsers := s.userRepository.FindAll()
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error()}
	}

	resData := resultUsers.Result.(models.Users)

	return response.Response{Success: true, Message: "User list", Data: resData}
}

func (s *userService) ReadById(Id int) response.Response {
	log.Print("[userService]...ReadById")

	resultUsers := s.userRepository.FindById(Id)
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error(), Data: Id}
	}

	resData := resultUsers.Result.(models.User)

	return response.Response{Success: true, Message: "User Detail", Data: resData}
}

func (s *userService) ReadByEmail(email string) response.Response {
	log.Print("[userService]...ReadById")

	resultUsers := s.userRepository.FindByEmail(email)
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error(), Data: email}
	}

	// Mapp interface to model
	resData := resultUsers.Result.(models.User)

	return response.Response{Success: true, Message: "User Detail", Data: resData}
}

func (s *userService) ReadByUsername(username string) response.Response {
	log.Print("[userService]...ReadById")

	resultUsers := s.userRepository.FindByUsername(username)
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error(), Data: username}
	}

	resData := resultUsers.Result.(models.User)

	return response.Response{Success: true, Message: "User Detail", Data: resData}
}

func (s *userService) ReadByIdWithRelation(Id int) response.Response {
	log.Print("[userService]... ReadByIdWithRelation")

	resultUsers := s.userRepository.FindByIdWithRelation(Id)
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error(), Data: Id}
	}

	resData := resultUsers.Result.(models.User)

	return response.Response{Success: true, Message: "User Detail", Data: resData}
}

func (s *userService) Update(request models.User) response.Response {
	log.Print("[userService]...Update")

	existUserResponse := s.ReadById(request.ID)

	if !existUserResponse.Success {
		return existUserResponse
	}

	existUser := existUserResponse.Data.(models.User)

	t := time.Now().UTC()
	timeNow := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	existUser.RoleId = request.RoleId
	existUser.TeamId = request.TeamId
	existUser.Name = request.Name
	existUser.Username = request.Username
	existUser.Email = request.Email
	existUser.EmailVerifiedAt = &timeNow

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return response.Response{Success: false, Message: err.Error()}
	}

	existUser.Password = string(passwordHash)

	resultUserUpdated := s.userRepository.Update(existUser)

	if resultUserUpdated.Error != nil {
		return response.Response{Success: false, Message: resultUserUpdated.Error.Error()}
	}

	resData := resultUserUpdated.Result.(models.User)

	resultUserDetail := s.userRepository.FindByIdWithRelation(resData.ID)

	if resultUserDetail.Error != nil {
		return response.Response{Success: false, Message: resultUserDetail.Error.Error()}
	}

	return response.Response{Success: true, Message: "User Updated", Data: resultUserDetail.Result}

}

func (s *userService) Delete(requestId request.RequestId) response.Response {
	log.Print("[userService]...Delete")

	existUserResponse := s.ReadById(requestId.ID)

	if !existUserResponse.Success {
		return existUserResponse
	}

	existUser := existUserResponse.Data.(models.User)

	resultUserDeleted := s.userRepository.SoftDelete(existUser)

	if resultUserDeleted.Error != nil {
		return response.Response{Success: false, Message: resultUserDeleted.Error.Error()}
	}

	resData := resultUserDeleted.Result.(models.User)

	return response.Response{Success: true, Message: "User data has been deleted", Data: resData}
}

func (s *userService) Trash() response.Response {
	resultUsers := s.userRepository.FindAllWithTrashed()
	if resultUsers.Error != nil {
		return response.Response{Success: false, Message: resultUsers.Error.Error()}
	}

	resData := resultUsers.Result.(models.Users)

	return response.Response{Success: true, Message: "User list Trash", Data: resData}
}

func (s *userService) Restore(Id int) response.Response {

	getUserTrash := s.userRepository.FindSingleTrashedById(Id)
	if getUserTrash.Error != nil {
		return response.Response{Success: false, Message: getUserTrash.Error.Error()}
	}

	resultUserTrash := s.userRepository.Restore(Id)

	if resultUserTrash.Error != nil {
		return response.Response{Success: false, Message: resultUserTrash.Error.Error()}
	}

	resultUser := s.userRepository.FindByIdWithRelation(Id)

	if resultUser.Error != nil {
		return response.Response{Success: false, Message: resultUser.Error.Error()}
	}

	resData := resultUser.Result.(models.User)

	return response.Response{Success: true, Message: "User data has been restore", Data: resData}
}
