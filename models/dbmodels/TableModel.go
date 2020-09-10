package dbmodels

type Table struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Status       int32      `json:"status"`
	GroupTabelID string     `json:"groupTabelId"`
	GroupTable   GroupTable `gorm:"foreignkey:id; association_foreignkey:groupTabelId; association_autoupdate:false;association_autocreate:false"`
	OrderID      int64      `json:"orderId"`
}

// TableName ...
func (t *Table) TableName() string {
	return "public.tabels"
}
