package dbmodels

import "time"

type Voucher struct {
	ID 		int64 `json:"id"`
	Description 	string `json:"description"`
	DiscType string `json:"discType"`
	Value 	int64 `json:"value"`
	DateStart time.Time `json:"dateStart"`
	DateEnd time.Time `json:"dateEnd"`
	MaxValue int64 `json:"maxValue"`
	MinPayment	int64 `json:"minPayment"`
	Code 		string `json:"code"`
}


// TableName ...
func (t *Voucher) TableName() string {
	return "public.voucher"
}
