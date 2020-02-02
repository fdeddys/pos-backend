package repository

import (
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
	// "github.com/astaxie/beego/logs"
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

	if err := db.Preload("Resto").Order("id desc").Limit(limit).Offset((page-1)*limit).Where("customer_Id = ?", req.CustomerId).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

// GetOrderById ...
func GetOrderById(id int64) (dbmodels.Order, error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.Order{}

	err := db.Preload("Resto").Where(" id = ?  ", id).First(&order).Error

	return order, err

}

// GetOrderDetailByOrderID ...
func GetOrderDetailByOrderID(orderID int64) []dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []dbmodels.OrderDetail

	db.Preload("MenuItem").Find(&orderDetails, " order_Id = ? and qty > 0 ", orderID)

	return orderDetails
}

// GetByRestoIDPage ...
func GetByRestoIDPage(req dto.OrderRequestDto, page int, limit int) ([]dbmodels.Order, error) {
	db := database.GetDbCon()

	var orders []dbmodels.Order

	if err := db.Preload("Resto").Order("id desc").Limit(limit).Offset((page-1)*limit).Where("resto_Id = ?", req.RestoId).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

// UpdatePayment ...
func UpdatePayment(orderID int64, statusPay string) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var order dbmodels.Order

	err := db.Model(&order).Where(" Id = ?", orderID).Update("IsPaid", statusPay).Error

	return err
}

// UpdateQty ...
func UpdateQty(orderDetailID int64, qty int) (dbmodels.OrderDetail, error) {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail
	// logs.Info("exec ", orderDetailID, "qty ", qty)
	err := db.Model(&orderDetail).Where("id = ?", orderDetailID).Update("qty", qty).Error

	if err != nil {
		return orderDetail, err
	}

	return orderDetail, err
}

// UpdateTotal ...
func UpdateTotal(orderID int64, total int64) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var order dbmodels.Order

	err := db.Model(&order).Where("id = ?", orderID).Update("total", total).Error

	return err
}

// GetOrderDetailByOrderDetailID ...
func GetOrderDetailByOrderDetailID(orderDetailID int64) dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail

	db.Preload("MenuItem").Find(&orderDetail, " Id = ?", orderDetailID)

	return orderDetail
}
