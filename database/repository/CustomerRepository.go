package repository

import (
	"fmt"
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
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

// GetCustomerFilterPaging ...
func GetCustomerFilterPaging(req dto.CustomerDto, page int, limit int) ([]dbmodels.Customer, int, error) {
	db := database.GetDbCon()

	var customers []dbmodels.Customer
	var total int

	if err := db.Limit(limit).Offset((page - 1) * limit).Order("id").Find(&customers).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		log.Println("<<< Error get data Customer by filter paging >>>")

		return customers, 0, err
	}
	fmt.Println("iterate Customer ", customers)

	return customers, total, nil
}
