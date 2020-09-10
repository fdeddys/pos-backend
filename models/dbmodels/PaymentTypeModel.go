package dbmodels

// PaymentType ...
type PaymentType struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Urut    int32  `json:"urut"`
	Status  int32  `json:"status"`
	RestoID int64  `json:"restoId"`
}

// TableName ...
func (t *PaymentType) TableName() string {
	return "public.payment_types"
}
