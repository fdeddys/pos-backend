package dto

type MenuItemDto struct {
	ID 		int64 `json:"id"`
	GroupID 		int64 `json:"groupId"`
	RestoID 		int64 `json:"restoId"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	ImgUrl string `json:"imgURL"`
	Price float64 `json:"price"`
	Status 	int `json:"status"`
}