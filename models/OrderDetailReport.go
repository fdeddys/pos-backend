package models

import "time"

type OrderDetailReport struct {
	ID      int64   `json:"id"`
	OrderId int64	`json:"order_id"`

	EMenuItem int64 `json:"e_menu_item"`

	Qty int	`json:"qty"`
	Price 	float64 `json:"price"`
	Status string `json:"status"`
	OrderNo string `json:"order_no"`
	OrderDate time.Time `json:"order_date"`
	Customer string `json:"customer"`
	IsPaid string `json:"is_paid"`
	IsPaidDesc string `json:"is_paid_desc"`
	IsComplete 		string `json:"is_complete"`
	Notes 			string `json:"notes"`
	MenuItem 		string `json:"menu_item"`
	GrandTotal		string `json:"grand_total"`
	OrderStatus 	int `json:"order_status"`
	StatusDesc		string `json:"status_desc"`

}
