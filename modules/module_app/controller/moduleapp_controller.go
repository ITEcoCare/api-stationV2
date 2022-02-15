package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/module_app"
	"api-station/request"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type moduleAppController struct {
	moduleAppService module_app.IService
}

func NewModuleAppController(moduleAppService module_app.IService) *moduleAppController {
	return &moduleAppController{moduleAppService}
}

func (uC *moduleAppController) Create(c *gin.Context) {
	log.Print("[moduleAppController]... Create")

	var request models.ModuleApp

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.moduleAppService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC moduleAppController) Read(c *gin.Context) {
	log.Print("[moduleAppController]...Read")

	res := uC.moduleAppService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC moduleAppController) ReadById(c *gin.Context) {
	log.Print("[moduleAppController]...Read By Id")

	userId, _ := strconv.Atoi(c.Param("id"))

	res := uC.moduleAppService.ReadById(userId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC moduleAppController) Update(c *gin.Context) {
	log.Print("[moduleAppController]... Update")

	var request models.ModuleApp

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.moduleAppService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC moduleAppController) Delete(c *gin.Context) {
	log.Print("[moduleAppController]... Delete")

	var request request.RequestIdModuleApp

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.moduleAppService.Delete(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC moduleAppController) Trash(c *gin.Context) {
	log.Print("[moduleAppController]... Trash")

	res := uC.moduleAppService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC moduleAppController) Restore(c *gin.Context) {
	log.Print("[moduleAppController]... Restore")

	var request request.RequestIdUser

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.moduleAppService.Restore(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
