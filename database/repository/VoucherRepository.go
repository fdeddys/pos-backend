package repository

import (
	"errors"
	"resto-be/database"
	"resto-be/database/dbmodels"
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