package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/team"
	"api-station/request"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type teamController struct {
	teamService team.IService
}

func NewTeamController(teamService team.IService) *teamController {
	return &teamController{teamService}
}

func (uC *teamController) Create(c *gin.Context) {
	log.Print("[teamController]... Create")

	var request models.Team

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.teamService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC teamController) Read(c *gin.Context) {
	log.Print("[teamController]...Read")

	res := uC.teamService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC teamController) ReadById(c *gin.Context) {
	log.Print("[teamController]...Read By Id")

	userId, _ := strconv.Atoi(c.Param("id"))

	res := uC.teamService.ReadById(userId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC teamController) Update(c *gin.Context) {
	log.Print("[teamController]... Update")

	var request models.Team

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.teamService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC teamController) Delete(c *gin.Context) {
	log.Print("[teamController]... Delete")

	var request request.RequestIdteam

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.teamService.Delete(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC teamController) Trash(c *gin.Context) {
	log.Print("[teamController]... Trash")

	res := uC.teamService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC teamController) Restore(c *gin.Context) {
	log.Print("[teamController]... Restore")

	var request request.RequestIdteam

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.teamService.Restore(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
