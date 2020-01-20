package repository

import (
	"fmt"
	"resto-be/database"
	"resto-be/database/dbmodels"
)

// GetCustomerByEmail ...
func GetCustomerByEmail(email string) (dbmodels.Customer, error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var customer dbmodels.Customer
	var err error

	db.Where("email = ?", email).Find(&customer)

	fmt.Println("Customer => ", customer)
	return customer, err

}

// SaveCustomer ...
func SaveCustomer(customer *dbmodels.Customer) error {
	db := database.GetDbCon()

	err := db.Save(&customer).Error

	return err
}
