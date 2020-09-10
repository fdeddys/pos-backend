package repository

import (
	"resto-be/database"
	"resto-be/models/dbmodels"
	"time"
)

func SaveOrderPayment(orderPayment *dbmodels.OrderPayment) error {
	db := database.GetDbCon()
	err := db.Save(&orderPayment).Error

	return err
}

// GetOrderPaymentByOrderId ...
func GetOrderPaymentByOrderId(orderID int64) []dbmodels.OrderPayment {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderpayments []dbmodels.OrderPayment

	db.Find(&orderpayments, " order_Id = ? ", orderID)

	return orderpayments

}

func UpdateAmountById(ID int64, amount float64) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	updateTime := time.Now()
	var orderPayment dbmodels.OrderPayment

	err := db.Model(&orderPayment).Where(" Id = ?", ID).Update(dbmodels.OrderPayment{Total: amount, UpdateDate: updateTime}).Error

	return err
}
