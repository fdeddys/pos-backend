package dto

type MenuItemRequestDto struct {
	ID 		int64 `json:"id"`
	GroupID 		int64 `json:"groupId"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	ImgUrl string `json:"imgURL"`
	Price float64 `json:"price"`
	Stock   int `json:"stock"`
	IsFavorite  int     `json:"isFavorite"`
	Status 	int `json:"status"`
	RestoId int64 `json:"restoId"`
	CategoryId int64 `json:"categoryId"`
}

type UploadImageMenuItemRequestDto struct {
	ID 		int64 `json:"id"`
	Data string `json:"data"`
	MenuItemId int64 `json:"menuItemId"`
}