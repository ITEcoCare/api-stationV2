package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/permission"
	"api-station/request"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type permissionController struct {
	permissionService permission.IService
}

func NewPermissionController(permissionService permission.IService) *permissionController {
	return &permissionController{permissionService}
}

func (uC *permissionController) Create(c *gin.Context) {
	log.Print("[permissionController]... Create")

	var request models.Permission

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.permissionService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC permissionController) Read(c *gin.Context) {
	log.Print("[permissionController]...Read")

	res := uC.permissionService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC permissionController) ReadById(c *gin.Context) {
	log.Print("[permissionController]...Read By Id")

	userId, _ := strconv.Atoi(c.Param("id"))

	res := uC.permissionService.ReadById(userId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC permissionController) Update(c *gin.Context) {
	log.Print("[permissionController]... Update")

	var request models.Permission

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.permissionService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC permissionController) Delete(c *gin.Context) {
	log.Print("[permissionController]... Delete")

	var request request.RequestIdPermission

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.permissionService.Delete(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC permissionController) Trash(c *gin.Context) {
	log.Print("[permissionController]... Trash")

	res := uC.permissionService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC permissionController) Restore(c *gin.Context) {
	log.Print("[permissionController]... Restore")

	var request request.RequestIdPermission

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.permissionService.Restore(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
