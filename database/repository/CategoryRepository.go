package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
)

func GetAllCategory() ([]dbmodels.Category, error) {
	db := database.GetDbCon()

	var categories []dbmodels.Category

	err := db.Find(&categories).Error

	return categories, err
}