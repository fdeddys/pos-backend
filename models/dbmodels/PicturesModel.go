package dbmodels

type RestoPicture struct {
	ID 		int64 `json:"id"`
	ImgUrl 	string `json:"imgURL"`
	RestoId int64 `json:"restoId"`
	Status 	int `json:"status"`
}
func (t *RestoPicture) TableName() string{
	return "public.resto_picture"
}


type MenuItemPicture struct {
	ID 		int64 `json:"id"`
	ImgUrl 	string `json:"imgURL"`
	MenuItemId int64 `json:"menuItemId"`
	Status 	int `json:"status"`
}
func (t *MenuItemPicture) TableName() string{
	return "public.menu_item_picture"
}