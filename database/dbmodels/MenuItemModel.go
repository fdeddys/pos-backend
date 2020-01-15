package dbmodels

type MenuItem struct {
	ID 		int64 `json:"id"`
	GroupID 		int64 `json:"group_id"`
	RestoID 		int64 `json:"resto_id"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	ImgUrl string `json:"img_url"`
	Price float64 `json:"price"`
	Status 	int `json:"status"`
}


// TableName ...
func (t *MenuItem) TableName() string {
	return "public.e_menu_item"
}