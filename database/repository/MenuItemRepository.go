package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
)

func SaveMenuItem(menuItem *dbmodels.MenuItem) (error) {
	db := database.GetDbCon()

	err := db.Save(&menuItem).Error

	return err
}