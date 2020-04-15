package dbmodels

type GroupTable struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int32  `json:"status"`
	RestoCode string `json:"restoCode"`
}


// TableName ...
func (t *GroupTable) TableName() string {
	return "public.group_tables"
}

