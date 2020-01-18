package dto

type MenuItemDto struct {
	ID 		int64 `json:"id"`
	GroupID 		int64 `json:"groupId"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	ImgUrl string `json:"imgURL"`
	Price float64 `json:"price"`
	Stock   int `json:"stock"`
	Status 	int `json:"status"`
}