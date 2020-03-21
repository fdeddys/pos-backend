package dbmodels

type Category struct {
	ID 		int64 `json:"id"`
	Name	string `json:"name"`
}

// TableName ...
func (t *Category) TableName() string {
	return "public.category"
}