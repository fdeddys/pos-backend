package repository

import (
	"resto-be/database"
	"resto-be/models/dbmodels"
)

func SaveGroupTable(voucher *dbmodels.GroupTable) (error) {
	db := database.GetDbCon()

	err := db.Save(&voucher).Error

	return err
}