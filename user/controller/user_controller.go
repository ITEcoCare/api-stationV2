package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/request"
	"api-station/user"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.IService
}

func NewUserController(userService user.IService) *userController {
	return &userController{userService}
}

func (uC *userController) Create(c *gin.Context) {
	log.Print("[userController]... Create User")

	var request models.User

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.userService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC userController) Read(c *gin.Context) {
	log.Print("[userController]...Read Users")

	res := uC.userService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC userController) ReadById(c *gin.Context) {
	log.Print("[userController]...Read ReadUser By Id")

	userId, _ := strconv.Atoi(c.Param("id"))

	res := uC.userService.ReadByIdWithRelation(userId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC userController) Update(c *gin.Context) {
	log.Print("[userController]...UpdateUser")

	var request models.User

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.userService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC userController) Delete(c *gin.Context) {
	log.Print("[userController]...DeleteUser")

	var requestId request.RequestId

	err := c.ShouldBindJSON(&requestId)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.userService.Delete(requestId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC userController) Trash(c *gin.Context) {
	log.Print("[userController]... Trash")

	res := uC.userService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC userController) Restore(c *gin.Context) {
	log.Print("[userController]... Restore")

	var requestId request.RequestId

	err := c.ShouldBindJSON(&requestId)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.userService.Restore(requestId.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
