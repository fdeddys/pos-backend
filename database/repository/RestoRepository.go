package repository

import (
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
)

func SaveResto(resto *dbmodels.Resto) (error) {
	db := database.GetDbCon()

	err := db.Save(&resto).Error

	return err
}

func GetAllResto() ([]dbmodels.Resto, error) {
	db := database.GetDbCon()

	var restorants []dbmodels.Resto

	err := db.Find(&restorants).Error

	return restorants, err
}

func GetById(id int64) (dbmodels.Resto, error) {
	db := database.GetDbCon()

	var resto dbmodels.Resto

	err := db.Where(dbmodels.Resto{ID:id}).First(&resto).Error

	return resto, err
}

func GetRestoyFilterPaging(req dto.RestoRequesDto, page int, limit int) ([]dbmodels.Resto, int, error) {
	db := database.GetDbCon()

	var restorants []dbmodels.Resto

	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id").Find(&restorants).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		log.Println("<<< Error get data restoran by filter paging >>>")
		return restorants, 0, err
	}


	return restorants, total, err
}