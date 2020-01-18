package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
)

func AddOrder(order *dbmodels.Order) (error) {
	db := database.GetDbCon()
	err := db.Save(&order).Error

	return err
}


func AddOrderDetail(orderDetail *dbmodels.OrderDetail) (error) {
	db := database.GetDbCon()
	err := db.Save(&orderDetail).Error

	return err
}