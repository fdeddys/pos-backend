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

	if data.Email == "" || data.Password == ""{
		res.Rc = constants.ERR_CODE_41
		res.Msg = constants.ERR_CODE_41_MSG
		return res
	}

	c:= service.CekDataCustomer(data.Email)
	log.Println("c ==>", c)

	if c == false {
		res.Rc = constants.ERR_CODE_42
		res.Msg = constants.ERR_CODE_42_MSG
		return res
	}

	data.ManualRestoID = dto.CurrRestoID

	dataCustomer := dbmodels.Customer{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		PhoneNumb:      data.PhoneNumb,
		Fb:             data.Fb,
		Password:       data.Password,
		ManualCustomer: 1,
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

func (service *CustomerServiceInterface) CekDataCustomer(email string) bool{
	restoId := dto.CurrRestoID

	_, err:= repository.GetCustomerEmailAndRestoId(email,restoId)
	log.Println("err ==> ", err)
	if err == nil {
		return false
	}
	return true
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
