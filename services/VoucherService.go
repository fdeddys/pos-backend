package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
)

type VoucherServiceInterface struct {
}

func InitializeVoucherServiceInterface() *VoucherServiceInterface {
	return &VoucherServiceInterface{}
}

func (service *VoucherServiceInterface) GetByCode(code string) models.Response {

	var res models.Response

	voucher, err := repository.GetVoucherByCode(code)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = err.Error()
		res.Data = voucher
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = voucher

	return res
}