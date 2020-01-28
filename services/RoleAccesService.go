package services

import "resto-be/database/repository"

type RoleAccessService struct {
	
}

func InitRoleAccessService() *RoleAccessService  {
	return &RoleAccessService{}
}

func (service *RoleAccessService) GetUserAccess(userID int64, accessName string) bool {

	return repository.CountUserAccessByUserIdAndAccessName(userID, accessName)
}