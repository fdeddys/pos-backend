package services

import (
	"log"
	"resto-be/database/dbmodels"
	"resto-be/models"
	"resto-be/models/dto"
)

type MenuGroupInterface struct {

}

func InitializeMenuGroupInterface()  *MenuGroupInterface {
	return &MenuGroupInterface{
	}
}

func (service *MenuGroupInterface) Save (restoDto *dto.MenuGroupRequestDto) models.Response{
	var res models.Response

	menuGroup := dbmodels.MenuGroup{}

	log.Println(menuGroup)
	return res
}