package service

import (
	"api-station/models"
	"api-station/modules/permission"
	"api-station/response"
	"log"
)

type permissionService struct {
	permissionRepository permission.IRepository
}

func NewPermissionService(permissionRepository permission.IRepository) *permissionService {
	return &permissionService{permissionRepository}
}

func (s *permissionService) Create(request models.Permission) response.Response {
	log.Print("[permissionService]...Create")

	permission := models.Permission{}
	permission.Name = request.Name
	permission.Description = request.Description

	resultPermission := s.permissionRepository.Save(permission)

	if resultPermission.Error != nil {
		return response.Response{Success: false, Message: resultPermission.Error.Error()}
	}

	resData := resultPermission.Result.(models.Permission)

	return response.Response{Success: true, Message: "Permission has been created", Data: resData}
}

func (s *permissionService) Read() response.Response {
	log.Print("[permissionService]... Read")

	resultPermissions := s.permissionRepository.FindAll()
	if resultPermissions.Error != nil {
		return response.Response{Success: false, Message: resultPermissions.Error.Error()}
	}

	resData := resultPermissions.Result.(models.Permissions)

	return response.Response{Success: true, Message: "Permission list", Data: resData}
}

func (s *permissionService) ReadById(Id int) response.Response {
	log.Print("[permissionService]...ReadById")

	resultPermission := s.permissionRepository.FindById(Id)
	if resultPermission.Error != nil {
		return response.Response{Success: false, Message: resultPermission.Error.Error(), Data: Id}
	}

	resData := resultPermission.Result.(models.Permission)

	return response.Response{Success: true, Message: "Permission Detail", Data: resData}
}

func (s *permissionService) Update(request models.Permission) response.Response {
	log.Print("[permissionService]... Update")

	resultPermissionExist := s.ReadById(request.ID)

	if !resultPermissionExist.Success {
		return resultPermissionExist
	}

	existPermission := resultPermissionExist.Data.(models.Permission)

	existPermission.Name = request.Name
	existPermission.Description = request.Description

	resultPermissionUpdated := s.permissionRepository.Update(existPermission)

	if resultPermissionUpdated.Error != nil {
		return response.Response{Success: false, Message: resultPermissionUpdated.Error.Error()}
	}

	resData := resultPermissionUpdated.Result.(models.Permission)

	return response.Response{Success: true, Message: "Permission Updated", Data: resData}

}

func (s *permissionService) Delete(Id int) response.Response {
	log.Print("[permissionService]...Delete")

	resultPermissionExist := s.ReadById(Id)

	if !resultPermissionExist.Success {
		return resultPermissionExist
	}

	permission := resultPermissionExist.Data.(models.Permission)

	resultPermissionDeleted := s.permissionRepository.SoftDelete(permission)

	if resultPermissionDeleted.Error != nil {
		return response.Response{Success: false, Message: resultPermissionDeleted.Error.Error()}
	}

	resData := resultPermissionDeleted.Result.(models.Permission)

	return response.Response{Success: true, Message: "Permission has been deleted", Data: resData}
}

func (s *permissionService) Trash() response.Response {
	resultPermissionsExist := s.permissionRepository.FindAllWithTrashed()
	if resultPermissionsExist.Error != nil {
		return response.Response{Success: false, Message: resultPermissionsExist.Error.Error()}
	}

	resData := resultPermissionsExist.Result.(models.Permissions)

	return response.Response{Success: true, Message: "Permission list Trash", Data: resData}
}

func (s *permissionService) Restore(Id int) response.Response {

	resultRoleExist := s.permissionRepository.FindSingleTrashedById(Id)
	if resultRoleExist.Error != nil {
		return response.Response{Success: false, Message: resultRoleExist.Error.Error()}
	}

	resultRoleRestore := s.permissionRepository.Restore(Id)

	if resultRoleRestore.Error != nil {
		return response.Response{Success: false, Message: resultRoleRestore.Error.Error()}
	}

	return response.Response{Success: true, Message: "Permission data has been restore", Data: Id}
}
