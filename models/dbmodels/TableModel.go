package dbmodels

type Table struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int32  `json:"status"`
	GroupTabelID string `json:"groupTabelId"`
}


// TableName ...
func (t *Table) TableName() string {
	return "public.tabels"
}

