package repository

import (
	"errors"
	"log"
	"resto-be/database"
	"resto-be/database/dbmodels"
	"resto-be/models/dto"
	"time"
)

func GetVoucherByCode(code string) (dbmodels.Voucher, error) {
	db := database.GetDbCon()

	var voucher dbmodels.Voucher

	err := db.Where("code = ?", code).First(&voucher).Error

	if err != nil {
		return voucher, errors.New("Kode Voucher Tidak Ditemukan")
	}
	now := time.Now()


	if now.Before(voucher.DateStart) {
		return voucher, errors.New("Kode Voucher Belum Aktif")
	}

	if now.After(voucher.DateEnd) {
		return voucher, errors.New("Kode Voucher Kadaluarsa")
	}


	return voucher, err
}

func SaveVoucher(voucher *dbmodels.Voucher) (error) {
	db := database.GetDbCon()

	err := db.Save(&voucher).Error

	return err
}


func GetVoucherById(id int64) (dbmodels.Voucher, error) {
	db := database.GetDbCon()

	var voucher dbmodels.Voucher

	if id == 0 {
		return voucher, errors.New("id = 0")
	}

	err := db.Where(dbmodels.Voucher{ID:id}).First(&voucher).Error

	return voucher, err
}

func GetVoucherFilterPaging(req dto.VoucherRequestDto, page int, limit int) ([]dbmodels.Voucher, int, error) {
	db := database.GetDbCon()

	var vouchers []dbmodels.Voucher

	var total int


	err := db.Where(dbmodels.Voucher{
	}).Limit(limit).Offset((page-1) * limit).Order("id").Find(&vouchers).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		log.Println("<<< Error get data voucher by filter paging >>>")
		return vouchers, 0, err
	}


	return vouchers, total, err
}