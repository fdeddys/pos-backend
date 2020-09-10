package dto

import "time"

type OrderPaymentDto struct {
	ID              int64     `json:"id"`
	OrderID         int64     `json:"orderId"`
	PaymentTypeID   int64     `json:"paymentTypeId"`
	UpdateDate      time.Time `json:"orderDate"`
	PaymentTypeName string    `json:"paymentTypeName"`
	Total           float64   `json:"total"`
}
