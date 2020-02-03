package services

import (
	"log"
	"resto-be/constants"
	dbmodels "resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
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

	data.ManualRestoID = dto.CurrRestoID

	dataCustomer := dbmodels.Customer{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		PhoneNumb:      data.PhoneNumb,
		Fb:             data.Fb,
		Password:       data.Password,
		ManualCustomer: data.ManualCustomer,
		ManualRestoID:  data.ManualRestoID,
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

// GetDataCustomerByFilterPaging ...
func (service *CustomerServiceInterface) GetDataCustomerByFilterPaging(req dto.CustomerDto, page int, count int) models.Response {
	var res models.Response

	customers, total, err := repository.GetCustomerFilterPaging(req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = customers
	res.TotalData = total

	return res

}

// GetCustByID ...
func (service *CustomerServiceInterface) GetCustByID(id int64) models.Response {
	var res models.Response

	resto, err := repository.GetCustomerByID(id)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto

	return res

}
