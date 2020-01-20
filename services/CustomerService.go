package services

import (
	"log"
	"resto-be/constants"
	dbmodels "resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
)

// CustomerServiceInterface ...
type CustomerServiceInterface struct {
}

// InitializeCustomerServiceInterface ...
func InitializeCustomerServiceInterface() *CustomerServiceInterface {
	return &CustomerServiceInterface{}
}

// SaveDataCustomer ...
func (service *CustomerServiceInterface) SaveDataCustomer(data *dbmodels.Customer) models.Response {
	var res models.Response

	dataCustomer := dbmodels.Customer{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		PhoneNumb: data.PhoneNumb,
		Fb:        data.Fb,
		Password:  data.Password,
	}

	err := repository.SaveCustomer(&dataCustomer)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG + " " + err.Error()
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res

}
