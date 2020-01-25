package dto

type MenuItemRequestDto struct {
	ID 		int64 `json:"id"`
	GroupID 		int64 `json:"groupId"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	ImgUrl string `json:"imgURL"`
	Price float64 `json:"price"`
	Stock   int `json:"stock"`
	Status 	int `json:"status"`
	RestoId int64 `json:"restoId"`
}

type UploadImageMenuItemRequestDto struct {
	Data string `json:"data"`
	MenuItemId int64 `json:"menuItemId"`
}