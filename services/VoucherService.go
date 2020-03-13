package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/utils"
)

type VoucherServiceInterface struct {
}

func InitializeVoucherServiceInterface() *VoucherServiceInterface {
	return &VoucherServiceInterface{}
}

func (service *VoucherServiceInterface) Save (voucherDto *dto.VoucherRequestDto) models.Response{
	var res models.Response

	voucher := dbmodels.Voucher{
		ID: voucherDto.ID,
		DateStart: utils.ConvertStringToTime(voucherDto.DateStart),
		DateEnd: utils.ConvertStringToTime(voucherDto.DateEnd),
		Code: voucherDto.Code,
		Value: voucherDto.Value,
		Description: voucherDto.Description,
		DiscType: voucherDto.DiscType,
		MaxValue: voucherDto.MaxValue,
		MinPayment: voucherDto.MinPayment,
	}

	err := repository.SaveVoucher(&voucher)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = voucher


	log.Println(voucher)
	return res
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


func (service *VoucherServiceInterface) GetById (id int64) models.Response{
	var res models.Response

	voucher, err := repository.GetVoucherById(id)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = voucher

	return res

}

func (service *VoucherServiceInterface) GetDataByFilterPaging (req dto.VoucherRequestDto, page int, count int) models.Response{
	var res models.Response

	vouchers, total, err := repository.GetVoucherFilterPaging(req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = vouchers
	res.TotalData = total

	return res

}