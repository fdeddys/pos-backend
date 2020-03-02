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
	User       User      `gorm:"foreignkey:id; association_foreignkey:UserId; association_autoupdate:false;association_autocreate:false"`
	Total      int64     `json:"total"`
	Status     int64     `json:"status"`
	StatusDesc     string     `json:"statusDesc"gorm:"-"`
	IsPaid     string     `json:"isPaid"`
	IsPaidDesc string		`json:"isPaidDesc"gorm:"-"`
	IsComplete string	`json:"isComplete"`
	IsCompleteDesc string	`json:"isCompleteDesc" gorm:"-"`
	CustomerId int64     `json:"customerId"`
	Customer       Customer      `gorm:"foreignkey:id; association_foreignkey:CustomerId; association_autoupdate:false;association_autocreate:false"`
	Notes      string    `json:"notes"`
}

func (t *Order) TableName() string {
	return "public.order"
}
