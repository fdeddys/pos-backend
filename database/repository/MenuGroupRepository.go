package repository

import (
	"errors"
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
)

func SaveMenuGroup(menuGroup *dbmodels.MenuGroup) (error) {
	db := database.GetDbCon()

	err := db.Save(&menuGroup).Error

	return err
}

func GetAllMenuGroup() ([]dbmodels.MenuGroup, error) {
	db := database.GetDbCon()

	var menuGroups []dbmodels.MenuGroup

	err := db.Find(&menuGroups).Error

	return menuGroups, err
}

func GetMenuGroupById(id int64) (dbmodels.MenuGroup, error) {
	db := database.GetDbCon()

	var menuGroup dbmodels.MenuGroup

	if id == 0 {
		return menuGroup, errors.New("id = 0")
	}

	err := db.Where(dbmodels.MenuGroup{ID:id}).First(&menuGroup).Error

	return menuGroup, err
}

func GetMenuGroupFilterPaging(req dto.MenuGroupRequestDto, page int, limit int) ([]dbmodels.MenuGroup, int, error) {
	db := database.GetDbCon()

	var menuGroups []dbmodels.MenuGroup

	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id").Find(&menuGroups).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		log.Println("<<< Error get data menuGroups by filter paging >>>")
		return menuGroups, 0, err
	}


	return menuGroups, total, err
}