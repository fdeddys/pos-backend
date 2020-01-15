package dbmodels

type MenuGroup struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	ImgUrl 	string `json:"img_url"`
	Status 	int `json:"status"`
}

// TableName ...
func (t *MenuGroup) TableName() string {
	return "public.e_menu_group"
}