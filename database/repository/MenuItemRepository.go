package repository

import (
	"errors"
	"resto-be/database"
	"resto-be/database/dbmodels"
)

func SaveMenuItem(menuItem *dbmodels.MenuItem) (error) {
	db := database.GetDbCon()

	err := db.Save(&menuItem).Error

	return err
}

func GetAllMenuItem() ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	err := db.Find(&menuItems).Error

	return menuItems, err
}

func GetMenuItemById(id int64) (dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItem dbmodels.MenuItem

	if id == 0 {
		return menuItem, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{ID:id}).First(&menuItem).Error
	return menuItem, err
}

func GetMenuItemByMenuGroupId(id int64) ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	if id == 0 {
		return menuItems, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{GroupID:id}).Find(&menuItems).Error
	return menuItems, err
}

func GetMenuItemByMenuGroupIdAndRestoId(groupId int64, restoId int64) ([]dbmodels.MenuItem, error) {
	db := database.GetDbCon()

	var menuItems []dbmodels.MenuItem

	if groupId == 0 || restoId == 0{
		return menuItems, errors.New("id = 0")
	}
	err := db.Where(dbmodels.MenuItem{GroupID:groupId, RestoID:restoId}).Find(&menuItems).Error
	return menuItems, err
}
