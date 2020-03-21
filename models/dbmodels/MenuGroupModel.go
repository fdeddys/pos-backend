package dbmodels

type MenuGroup struct {
	ID     int64  `json:"id"`
	RestoId int64 `json:"restoId"`
	Name   string `json:"name"`
	ImgUrl string `json:"imgURL"`
	Status int    `json:"status"`
}

// TableName ...
func (t *MenuGroup) TableName() string {
	return "public.e_menu_group"
}
