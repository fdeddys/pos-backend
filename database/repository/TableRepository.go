package repository

import (
	"log"
	"resto-be/database"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

func SaveTable(voucher *dbmodels.Table) error {
	db := database.GetDbCon()

	err := db.Save(&voucher).Error

	return err
}

func GetTableFilter(req dto.TableRequestDto) ([]dbmodels.Table, error) {
	db := database.GetDbCon()

	var groupTables []dbmodels.Table

	err := db.Preload("GroupTable").Table("tabels a").Select(" a.*").Joins("join group_tables b on a.group_tabel_id = b.id").Where("b.resto_code = ?", req.RestoCode).Find(&groupTables).Error // query

	if err != nil {
		log.Println("<<< Error get data grouptable by filter >>>")
		return groupTables, err
	}

	return groupTables, err
}

//GetTabelByTableID ...
func GetTabelByTableID(tableID int64) (dbmodels.Table, error) {

	db := database.GetDbCon()

	var tabel dbmodels.Table
	err := db.Preload("GroupTable").Where("id = ?", tableID).Find(&tabel).Error // query

	if err != nil {
		return tabel, err
	}

	return tabel, nil

}

func UpdateOrderIdToTabel(tabelID, orderID int64) error {

	// 20 - status Occupied
	return updateOrderIdAndStatusByTabelID(tabelID, orderID, 20)
}

func ReleaseTabel(tabelID int64) error {
	// 10 - status tabel empty
	return updateOrderIdAndStatusByTabelID(tabelID, 0, 10)
}

func updateOrderIdAndStatusByTabelID(tabelID, orderID int64, statusCode int) error {
	db := database.GetDbCon()

	var tabel dbmodels.Table
	err := db.Model(&tabel).Where("Id = ?", tabelID).Update("order_id", orderID).Update("status", statusCode).Error

	return err
}

func GetOrderIdFromTabelID(tabelID int64) (int64, error) {

	db := database.GetDbCon()
	db.Debug().LogMode(true)
	tabel := dbmodels.Table{}

	err := db.Where("id = ?  ", tabelID).First(&tabel).Error

	if err != nil {
		return 0, err
	}
	return tabel.OrderID, nil

}
