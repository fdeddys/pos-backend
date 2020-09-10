package dbmodels

import "time"

type OrderPayment struct {
	ID              int64     `json:"id"`
	OrderID         int64     `json:"orderId"`
	PaymentTypeId   int64     `json:"paymentTypeId"`
	UpdateDate      time.Time `json:"orderDate"`
	PaymentTypeName string    `json:"paymentTypeName"`
	Total           float64   `json:"total"`
}

func (t *OrderPayment) TableName() string {
	return "public.order_payment"
}
