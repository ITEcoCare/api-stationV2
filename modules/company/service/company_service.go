package service

import (
	"api-station/helpers"
	"api-station/models"
	"api-station/modules/company"
	"api-station/response"
	"fmt"
	"log"
	"strings"
)

type serviceCompany struct {
	repositoryCompany company.IRepository
}

func NewServiceCompany(repositoryCompany company.IRepository) *serviceCompany {
	return &serviceCompany{repositoryCompany}
}

func (s *serviceCompany) Create(request models.Company) response.Response {
	log.Print("[serviceCompany]...Create")

	if len(request.Code) > 10 {
		return response.Response{Success: false, Message: "Max CODE 10 Character"}
	}

	company := models.Company{}
	company.Code = strings.ToUpper(helpers.RemoveSpecialChar(strings.ReplaceAll(request.Code, " ", "")))
	company.Name = strings.Title(strings.ToLower(helpers.RemoveSpecialChar(request.Name)))
	company.Description = request.Description

	result := s.repositoryCompany.Save(company)

	if result.Error != nil {
		return response.Response{Success: false, Message: result.Error.Error()}
	}

	// map interface to models
	resData := result.Result.(models.Company)
	return response.Response{Success: true, Message: "Company registered", Data: resData}
}

func (s *serviceCompany) Read() response.Response {
	log.Print("[serviceCompany]...Read")

	result := s.repositoryCompany.FindAll()
	if result.Error != nil {
		return response.Response{Success: false, Message: result.Error.Error()}
	}

	resData := result.Result.(models.Companies)

	return response.Response{Success: true, Message: "Company list", Data: resData}
}

func (s *serviceCompany) ReadById(Id int) response.Response {
	log.Print("[serviceCompany]...ReadById")

	result := s.repositoryCompany.FindById(Id)
	if result.Error != nil {
		return response.Response{Success: false, Message: result.Error.Error(), Data: Id}
	}

	resData := result.Result.(models.Company)

	return response.Response{Success: true, Message: "Company Detail", Data: resData}
}

func (s *serviceCompany) Update(request models.Company) response.Response {
	log.Print("[serviceCompany]...Update")

	result := s.ReadById(request.ID)

	if !result.Success {
		return result
	}

	// map interface to model / dto
	company := result.Data.(models.Company)

	company.Code = strings.ToUpper(helpers.RemoveSpecialChar(strings.ReplaceAll(request.Code, " ", "")))
	company.Name = strings.Title(strings.ToLower(helpers.RemoveSpecialChar(request.Name)))
	company.Description = request.Description

	resultUpdated := s.repositoryCompany.Update(company)

	if resultUpdated.Error != nil {
		return response.Response{Success: false, Message: resultUpdated.Error.Error()}
	}

	resData := resultUpdated.Result.(models.Company)

	return response.Response{Success: true, Message: "Company Updated", Data: resData}

}

func (s *serviceCompany) Delete(Id int) response.Response {
	log.Print("[serviceCompany]...Delete")

	result := s.ReadById(Id)

	if !result.Success {
		return result
	}

	company := result.Data.(models.Company)

	resultDeleted := s.repositoryCompany.SoftDelete(company)

	if resultDeleted.Error != nil {
		return response.Response{Success: false, Message: resultDeleted.Error.Error()}
	}
	fmt.Println(resultDeleted)
	resData := resultDeleted.Result.(models.Company)

	return response.Response{Success: true, Message: "Company deleted", Data: resData}
}

func (s *serviceCompany) Trash() response.Response {
	result := s.repositoryCompany.FindAllWithTrashed()
	if result.Error != nil {
		return response.Response{Success: false, Message: result.Error.Error()}
	}

	resData := result.Result.(models.Companies)

	return response.Response{Success: true, Message: "Company list Trash", Data: resData}
}

func (s *serviceCompany) Restore(Id int) response.Response {

	result := s.repositoryCompany.FindSingleTrashedById(Id)
	if result.Error != nil {
		return response.Response{Success: false, Message: result.Error.Error()}
	}

	resultRestore := s.repositoryCompany.Restore(Id)

	if resultRestore.Error != nil {
		return response.Response{Success: false, Message: resultRestore.Error.Error()}
	}

	return response.Response{Success: true, Message: "Company data has been restore", Data: Id}
}
