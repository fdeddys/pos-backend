package dbmodels

type OrderDetail struct {
	ID      int64   `json:"id"`
	OrderId int64	`json:"orderId"`
	
	EMenuItem int64 `json:"eMenuItem"`
	MenuItem      MenuItem     `gorm:"foreignkey:id; association_foreignkey:EMenuItem; association_autoupdate:false;association_autocreate:false"`

	Qty int	`json:"qty"`
	Price 	float64 `json:"price"`
	Status string `json:"status"`
	StatusDesc string `json:"statusDesc"gorm:"-"`
}

func (t *OrderDetail) TableName() string{
	return "public.order_detail"
}
