package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
)

func SaveResto(resto *dbmodels.Resto) (error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)


	err := db.Save(&resto).Error

	return err
}