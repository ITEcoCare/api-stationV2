package service

import (
	"api-station/models"
	"api-station/modules/team"
	"api-station/request"
	"api-station/response"
	"log"
)

type teamService struct {
	teamRepository team.IRepository
}

func NewTeamService(teamRepository team.IRepository) *teamService {
	return &teamService{teamRepository}
}

func (s *teamService) Create(request models.Team) response.Response {
	log.Print("[teamService]...Create")

	team := models.Team{}
	team.Name = request.Name
	team.Description = request.Description

	resultTeam := s.teamRepository.Save(team)

	if resultTeam.Error != nil {
		return response.Response{Success: false, Message: resultTeam.Error.Error()}
	}

	newTeam := resultTeam.Result.(models.Team)

	return response.Response{Success: true, Message: "Team has been created", Data: newTeam}
}

func (s *teamService) Read() response.Response {
	log.Print("[teamService]...Read")

	resultTeams := s.teamRepository.FindAll()
	if resultTeams.Error != nil {
		return response.Response{Success: false, Message: resultTeams.Error.Error()}
	}

	resData := resultTeams.Result.(models.Teams)

	return response.Response{Success: true, Message: "Team list", Data: resData}
}

func (s *teamService) ReadById(Id int) response.Response {
	log.Print("[teamService]...ReadById")

	resultTeam := s.teamRepository.FindById(Id)
	if resultTeam.Error != nil {
		return response.Response{Success: false, Message: resultTeam.Error.Error(), Data: Id}
	}

	resData := resultTeam.Result.(models.Team)

	return response.Response{Success: true, Message: "User Detail", Data: resData}
}

func (s *teamService) Update(request models.Team) response.Response {
	log.Print("[teamService]...Update")

	resultTeamExist := s.ReadById(request.ID)

	if !resultTeamExist.Success {
		return resultTeamExist
	}

	existTeam := resultTeamExist.Data.(models.Team)

	existTeam.Name = request.Name
	existTeam.Description = request.Description

	resultUserUpdated := s.teamRepository.Update(existTeam)

	if resultUserUpdated.Error != nil {
		return response.Response{Success: false, Message: resultUserUpdated.Error.Error()}
	}

	resData := resultUserUpdated.Result.(models.Team)

	resultUserDetail := s.teamRepository.FindById(resData.ID)

	if resultUserDetail.Error != nil {
		return response.Response{Success: false, Message: resultUserDetail.Error.Error()}
	}

	return response.Response{Success: true, Message: "Team Updated", Data: resultUserDetail.Result}

}

func (s *teamService) Delete(requestId request.RequestId) response.Response {
	log.Print("[teamService]...Delete")

	resultTeamExist := s.ReadById(requestId.ID)

	if !resultTeamExist.Success {
		return resultTeamExist
	}

	resultTeam := resultTeamExist.Data.(models.Team)

	resultTeamDeleted := s.teamRepository.SoftDelete(resultTeam)

	if resultTeamDeleted.Error != nil {
		return response.Response{Success: false, Message: resultTeamDeleted.Error.Error()}
	}

	resData := resultTeamDeleted.Result.(models.Team)

	return response.Response{Success: true, Message: "Team has been deleted", Data: resData}
}

func (s *teamService) Trash() response.Response {
	resultTeamsExist := s.teamRepository.FindAllWithTrashed()
	if resultTeamsExist.Error != nil {
		return response.Response{Success: false, Message: resultTeamsExist.Error.Error()}
	}

	resultTeams := resultTeamsExist.Result.(models.Teams)

	return response.Response{Success: true, Message: "Team list Trash", Data: resultTeams}
}

func (s *teamService) Restore(Id int) response.Response {

	getUserTrash := s.teamRepository.FindSingleTrashedById(Id)
	if getUserTrash.Error != nil {
		return response.Response{Success: false, Message: getUserTrash.Error.Error()}
	}

	resultUserTrash := s.teamRepository.Restore(Id)

	if resultUserTrash.Error != nil {
		return response.Response{Success: false, Message: resultUserTrash.Error.Error()}
	}

	resultUser := s.teamRepository.FindById(Id)

	if resultUser.Error != nil {
		return response.Response{Success: false, Message: resultUser.Error.Error()}
	}

	resData := resultUser.Result.(models.Team)

	return response.Response{Success: true, Message: "Team data has been restore", Data: resData}
}
