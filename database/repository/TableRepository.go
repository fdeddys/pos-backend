package repository

import (
	"log"
	"resto-be/database"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

func SaveTable(voucher *dbmodels.Table) (error) {
	db := database.GetDbCon()

	err := db.Save(&voucher).Error

	return err
}

func GetTableFilter(req dto.TableRequestDto,) ([]dbmodels.Table, error) {
	db := database.GetDbCon()

	var groupTables []dbmodels.Table

	err := db.Table("tabels a").Select(" a.*").Joins("join group_tables b on a.group_tabel_id = b.id").Where("b.resto_code = ?", req.RestoCode).Find(&groupTables).Error // query

	if err != nil {
		log.Println("<<< Error get data grouptable by filter >>>")
		return groupTables, err
	}


	return groupTables, err
}