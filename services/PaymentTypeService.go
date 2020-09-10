package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

// PaymentTypeService ...
type PaymentTypeService struct {
}

func InitPaymentTypeService() *PaymentTypeService {
	return &PaymentTypeService{}
}

func (service *PaymentTypeService) Save(req *dto.PaymentTypeRequestDto) models.Response {
	var res models.Response

	resto, errResto := repository.GetRestoById(dto.CurrRestoID)
	if errResto != nil {
		res.Rc = constants.ERR_CODE_20
		res.Msg = constants.ERR_CODE_20_MSG
		return res
	}

	paymentType := dbmodels.PaymentType{
		ID:      req.ID,
		Name:    req.Name,
		Status:  req.Status,
		Urut:    req.Urut,
		RestoID: resto.ID,
	}

	err := repository.SavePaymentType(&paymentType)
	if err != nil {
		log.Println("err save payment type : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = paymentType

	log.Println(paymentType)
	return res
}

func (service *PaymentTypeService) GetAllPaymentType(req dto.PaymentTypeRequestDto) models.Response {
	var res models.Response

	paymentTypes, err := repository.GetPaymentTypeByRestoID(req.RestoID)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = paymentTypes

	return res

}
