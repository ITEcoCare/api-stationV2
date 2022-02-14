package auth

import (
	"api-station/auth"
	"api-station/helpers"
	"api-station/models"
	"api-station/request"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService auth.IService
}

func NewAuthController(authService auth.IService) *authController {
	return &authController{authService}
}

func (aC authController) Login(c *gin.Context) {
	log.Print("[authController]... Login User")

	var requestLogin request.RequestLogin

	err := c.ShouldBindJSON(&requestLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ValidationInputResponse(err))
		return
	}

	resultUser := aC.authService.CheckLogin(requestLogin)

	code := http.StatusOK

	if !resultUser.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, resultUser)
}

func (aC *authController) Register(c *gin.Context) {
	log.Print("[authController]... Register")

	var request models.User

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := aC.authService.Register(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (aC *authController) CheckEmailAvailability(c *gin.Context) {
	log.Print("[authController]... CheckEmailAvailability")

	var request request.CheckEmailRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := aC.authService.CheckEmailAvailability(request.Email)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (aC *authController) CheckUsernameAvailability(c *gin.Context) {
	log.Print("[authController]... CheckUsernameAvailability")

	var request request.CheckUsernameRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := aC.authService.CheckUsernameAvailability(request.Username)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
