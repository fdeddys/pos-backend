package dbmodels

import "time"

type Order struct {
	ID         int64     `json:"id"`
	OrderNo    string    `json:"orderNo"`
	TableId    int64     `json:"tableId"`
	OrderDate  time.Time `json:"orderDate"`
	RestoId    int64     `json:"restoId"`
	Resto      Resto     `gorm:"foreignkey:id; association_foreignkey:RestoId; association_autoupdate:false;association_autocreate:false"`
	UserId     int64     `json:"userId"`
	Total      int64     `json:"total"`
	Status     int64     `json:"status"`
	IsPaid     int64     `json:"isPaid"`
	CustomerId int64     `json:"customerId"`
	Notes      string    `json:"notes"`
}

func (t *Order) TableName() string {
	return "public.order"
}
