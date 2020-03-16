package dto

type VoucherRequestDto struct {
	ID 		int64 `json:"id"`
	Description 	string `json:"description"`
	DiscType string `json:"discType"`
	Value 	int64 `json:"value"`
	DateStart string `json:"dateStart"`
	DateEnd string `json:"dateEnd"`
	MaxValue int64 `json:"maxValue"`
	MinPayment	int64 `json:"minPayment"`
	Code 		string `json:"code"`
	RestoId	int64 `json:"restoId"`
}
