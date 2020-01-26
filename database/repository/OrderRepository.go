package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
)

func AddOrder(order *dbmodels.Order) error {
	db := database.GetDbCon()
	err := db.Save(&order).Error

	return err
}

func AddOrderDetail(orderDetail *dbmodels.OrderDetail) error {
	db := database.GetDbCon()
	err := db.Save(&orderDetail).Error

	return err
}

func GetByCustomerIdPage(req dto.OrderRequestDto, page int, limit int) ([]dbmodels.Order, error) {
	db := database.GetDbCon()

	var orders []dbmodels.Order

	if err := db.Order("id desc").Limit(limit).Offset((page-1)*limit).Where("customer_Id = ?", req.CustomerId).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}
