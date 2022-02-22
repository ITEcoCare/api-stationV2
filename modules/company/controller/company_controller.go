package controller

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/company"
	"api-station/request"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type companyController struct {
	companyService company.IService
}

func NewCompanyController(companyService company.IService) *companyController {
	return &companyController{companyService}
}

func (uC *companyController) Create(c *gin.Context) {
	log.Print("[companyController]... Create")

	var request models.Company

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.companyService.Create(request)

	code := http.StatusCreated
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC companyController) Read(c *gin.Context) {
	log.Print("[companyController]... Read")

	res := uC.companyService.Read()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC companyController) ReadById(c *gin.Context) {
	log.Print("[companyController]...Read ReadCompany By Id")

	CompanyId, _ := strconv.Atoi(c.Param("id"))

	res := uC.companyService.ReadById(CompanyId)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC companyController) Update(c *gin.Context) {
	log.Print("[companyController]...Update")

	var request models.Company

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.companyService.Update(request)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)

}

func (uC companyController) Delete(c *gin.Context) {
	log.Print("[companyController]...DeleteCompany")

	var request request.RequestIdCompany
	fmt.Println(request)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.companyService.Delete(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC companyController) Trash(c *gin.Context) {
	log.Print("[companyController]... Trash")

	res := uC.companyService.Trash()

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}

func (uC companyController) Restore(c *gin.Context) {
	log.Print("[companyController]... Restore")

	var request request.RequestIdCompany

	err := c.ShouldBindJSON(&request)
	if err != nil {
		res := helpers.ValidationInputResponse(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := uC.companyService.Restore(request.ID)

	code := http.StatusOK
	if !res.Success {
		code = http.StatusBadRequest
	}

	c.JSON(code, res)
}
