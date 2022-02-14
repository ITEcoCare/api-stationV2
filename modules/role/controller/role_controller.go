package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/role"
	"api-station/request"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleService role.IService
}

func NewRoleController(roleService role.IService) *roleController {
	return &roleController{roleService}
}

func (uC *roleController) Create(c *gin.Context) {
	log.Print("[roleController]... Create")

	var request models.Role

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.roleService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC roleController) Read(c *gin.Context) {
	log.Print("[roleController]...Read")

	res := uC.roleService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC roleController) ReadById(c *gin.Context) {
	log.Print("[roleController]...Read By Id")

	userId, _ := strconv.Atoi(c.Param("id"))

	res := uC.roleService.ReadById(userId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC roleController) Update(c *gin.Context) {
	log.Print("[roleController]... Update")

	var request models.Role

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.roleService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC roleController) Delete(c *gin.Context) {
	log.Print("[roleController]... Delete")

	var requestId request.RequestId

	err := c.ShouldBindJSON(&requestId)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.roleService.Delete(requestId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC roleController) Trash(c *gin.Context) {
	log.Print("[roleController]... Trash")

	res := uC.roleService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC roleController) Restore(c *gin.Context) {
	log.Print("[roleController]... Restore")

	var requestId request.RequestId

	err := c.ShouldBindJSON(&requestId)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.roleService.Restore(requestId.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
