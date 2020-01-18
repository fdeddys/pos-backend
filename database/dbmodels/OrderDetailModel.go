package dbmodels

type OrderDetail struct {
	ID      int64   `json:"id"`
	OrderId int64	`json:"orderId"`
	EMenuItem int64 `json:"eMenuItem"`
	Qty int	`json:"qty"`
	Price 	float64 `json:"price"`
}

func (t *OrderDetail) TableName() string{
	return "public.order_detail"
}
