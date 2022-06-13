package service

import (
	"api-station/models"
	"api-station/modules/module_app"
	"api-station/response"
	"log"
)

type moduleAppService struct {
	moduleAppRepository module_app.IRepository
}

func NewModuleAppService(moduleAppRepository module_app.IRepository) *moduleAppService {
	return &moduleAppService{moduleAppRepository}
}

func (s *moduleAppService) Create(request models.ModuleApp) response.Response {
	log.Print("[moduleAppService]...Create")

	moduleApp := models.ModuleApp{}
	moduleApp.Name = request.Name
	moduleApp.Description = request.Description

	resultModuleApp := s.moduleAppRepository.Save(moduleApp)

	if resultModuleApp.Error != nil {
		return response.Response{Success: false, Message: resultModuleApp.Error.Error()}
	}

	resData := resultModuleApp.Result.(models.ModuleApp)

	return response.Response{Success: true, Message: "Module App has been created", Data: resData}
}

func (s *moduleAppService) Read() response.Response {
	log.Print("[moduleAppService]...Read")

	resultModuleApps := s.moduleAppRepository.FindAll()
	if resultModuleApps.Error != nil {
		return response.Response{Success: false, Message: resultModuleApps.Error.Error()}
	}

	resData := resultModuleApps.Result.(models.ModuleApps)

	return response.Response{Success: true, Message: "Module App list", Data: resData}
}

func (s *moduleAppService) ReadById(Id int) response.Response {
	log.Print("[moduleAppService]...ReadById")

	resultModuleApp := s.moduleAppRepository.FindById(Id)
	if resultModuleApp.Error != nil {
		return response.Response{Success: false, Message: resultModuleApp.Error.Error(), Data: Id}
	}

	resData := resultModuleApp.Result.(models.ModuleApp)

	return response.Response{Success: true, Message: "Module App Detail", Data: resData}
}

func (s *moduleAppService) Update(request models.ModuleApp) response.Response {
	log.Print("[moduleAppService]... Update")

	resultModuleExist := s.ReadById(request.ID)

	if !resultModuleExist.Success {
		return resultModuleExist
	}

	existModule := resultModuleExist.Data.(models.ModuleApp)

	existModule.Name = request.Name
	existModule.Description = request.Description

	resultModuleUpdated := s.moduleAppRepository.Update(existModule)

	if resultModuleUpdated.Error != nil {
		return response.Response{Success: false, Message: resultModuleUpdated.Error.Error()}
	}

	resData := resultModuleUpdated.Result.(models.ModuleApp)

	return response.Response{Success: true, Message: "Module App Updated", Data: resData}

}

func (s *moduleAppService) Delete(Id int) response.Response {
	log.Print("[moduleAppService]...Delete")

	resultModuleExist := s.ReadById(Id)

	if !resultModuleExist.Success {
		return resultModuleExist
	}

	module := resultModuleExist.Data.(models.ModuleApp)

	resultModuleDeleted := s.moduleAppRepository.SoftDelete(module)

	if resultModuleDeleted.Error != nil {
		return response.Response{Success: false, Message: resultModuleDeleted.Error.Error()}
	}

	resData := resultModuleDeleted.Result.(models.ModuleApp)

	return response.Response{Success: true, Message: "Module App has been deleted", Data: resData}
}

func (s *moduleAppService) Trash() response.Response {
	resultModulesExist := s.moduleAppRepository.FindAllWithTrashed()
	if resultModulesExist.Error != nil {
		return response.Response{Success: false, Message: resultModulesExist.Error.Error()}
	}

	resData := resultModulesExist.Result.(models.ModuleApps)

	return response.Response{Success: true, Message: "Module App list Trash", Data: resData}
}

func (s *moduleAppService) Restore(Id int) response.Response {

	resultModuleExist := s.moduleAppRepository.FindSingleTrashedById(Id)
	if resultModuleExist.Error != nil {
		return response.Response{Success: false, Message: resultModuleExist.Error.Error()}
	}

	resultModuleRestore := s.moduleAppRepository.Restore(Id)

	if resultModuleRestore.Error != nil {
		return response.Response{Success: false, Message: resultModuleRestore.Error.Error()}
	}

	return response.Response{Success: true, Message: "Module App data has been restore", Data: Id}
}
