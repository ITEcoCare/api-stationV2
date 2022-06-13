package service

import (
	"api-station/models"
	"api-station/modules/role"
	"api-station/response"
	"log"
)

type roleService struct {
	roleRepository role.IRepository
}

func NewRoleService(roleRepository role.IRepository) *roleService {
	return &roleService{roleRepository}
}

func (s *roleService) Create(request models.Role) response.Response {
	log.Print("[roleService]...Create")

	role := models.Role{}
	role.CompanyId = request.CompanyId
	role.Name = request.Name
	role.Description = request.Description

	resultRole := s.roleRepository.Save(role)

	if resultRole.Error != nil {
		return response.Response{Success: false, Message: resultRole.Error.Error()}
	}

	newRole := resultRole.Result.(models.Role)

	return response.Response{Success: true, Message: "Role has been created", Data: newRole}
}

func (s *roleService) Read() response.Response {
	log.Print("[roleService]...Read")

	resultRoles := s.roleRepository.FindAll()
	if resultRoles.Error != nil {
		return response.Response{Success: false, Message: resultRoles.Error.Error()}
	}

	resData := resultRoles.Result.(models.Roles)

	return response.Response{Success: true, Message: "Role list", Data: resData}
}

func (s *roleService) ReadById(Id int) response.Response {
	log.Print("[roleService]...ReadById")

	resultRole := s.roleRepository.FindById(Id)
	if resultRole.Error != nil {
		return response.Response{Success: false, Message: resultRole.Error.Error(), Data: Id}
	}

	resData := resultRole.Result.(models.Role)

	return response.Response{Success: true, Message: "Role Detail", Data: resData}
}

func (s *roleService) Update(request models.Role) response.Response {
	log.Print("[roleService]... Update")

	resultRoleExist := s.ReadById(request.ID)

	if !resultRoleExist.Success {
		return resultRoleExist
	}

	existRole := resultRoleExist.Data.(models.Role)

	existRole.Name = request.Name
	existRole.Description = request.Description

	resultRoleUpdated := s.roleRepository.Update(existRole)

	if resultRoleUpdated.Error != nil {
		return response.Response{Success: false, Message: resultRoleUpdated.Error.Error()}
	}

	resData := resultRoleUpdated.Result.(models.Role)

	return response.Response{Success: true, Message: "Role Updated", Data: resData}

}

func (s *roleService) Delete(Id int) response.Response {
	log.Print("[roleService]...Delete")

	resultRoleExist := s.ReadById(Id)

	if !resultRoleExist.Success {
		return resultRoleExist
	}

	role := resultRoleExist.Data.(models.Role)

	resultRoleDeleted := s.roleRepository.SoftDelete(role)

	if resultRoleDeleted.Error != nil {
		return response.Response{Success: false, Message: resultRoleDeleted.Error.Error()}
	}

	resData := resultRoleDeleted.Result.(models.Role)

	return response.Response{Success: true, Message: "Role has been deleted", Data: resData}
}

func (s *roleService) Trash() response.Response {
	resultRolesExist := s.roleRepository.FindAllWithTrashed()
	if resultRolesExist.Error != nil {
		return response.Response{Success: false, Message: resultRolesExist.Error.Error()}
	}

	resData := resultRolesExist.Result.(models.Roles)

	return response.Response{Success: true, Message: "Role list Trash", Data: resData}
}

func (s *roleService) Restore(Id int) response.Response {

	resultRoleExist := s.roleRepository.FindSingleTrashedById(Id)
	if resultRoleExist.Error != nil {
		return response.Response{Success: false, Message: resultRoleExist.Error.Error()}
	}

	resultRoleRestore := s.roleRepository.Restore(Id)

	if resultRoleRestore.Error != nil {
		return response.Response{Success: false, Message: resultRoleRestore.Error.Error()}
	}

	return response.Response{Success: true, Message: "Role data has been restore", Data: Id}
}
