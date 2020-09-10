package repository

import (
	"log"
	"resto-be/database"
	"resto-be/models/dbmodels"
)

// SavePaymentType ...
func SavePaymentType(paymentType *dbmodels.PaymentType) error {
	db := database.GetDbCon()

	err := db.Save(&paymentType).Error

	return err
}

// GetPaymentTypeByRestoID ...
func GetPaymentTypeByRestoID(restoID int64) ([]dbmodels.PaymentType, error) {
	db := database.GetDbCon()
	var paymentTypes []dbmodels.PaymentType

	err := db.Where("resto_id = ? and status = 1 ", restoID).Find(&paymentTypes).Error // query

	if err != nil {
		log.Println("<<< Error get all Payment type >>>")
		return paymentTypes, err
	}

	return paymentTypes, err
}
