package repository

import (
	"log"
	"resto-be/database"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

func SaveGroupTable(voucher *dbmodels.GroupTable) (error) {
	db := database.GetDbCon()

	err := db.Save(&voucher).Error

	return err
}

func GetGroupTablePaging(req dto.GroupTableRequestDto, page int, limit int) ([]dbmodels.GroupTable, int, error) {
	db := database.GetDbCon()

	var groupTables []dbmodels.GroupTable

	var total int


	err := db.Where("resto_code = ?", req.RestoCode).Limit(limit).Offset((page-1) * limit).Order("id").Find(&groupTables).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		log.Println("<<< Error get data grouptable by filter paging >>>")
		return groupTables, 0, err
	}


	return groupTables, total, err
}

func GetGroupTableFilter(req dto.GroupTableRequestDto) ([]dbmodels.GroupTable, error) {
	db := database.GetDbCon()

	var groupTables []dbmodels.GroupTable


	err := db.Where("resto_code = ? ", req.RestoCode).Order("id").Find(&groupTables).Error // query

	if err != nil {
		log.Println("<<< Error get data grouptable by filter >>>")
		return groupTables, err
	}


	return groupTables, err
}