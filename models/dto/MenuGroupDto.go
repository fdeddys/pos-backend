package dto

type MenuGroupRequestDto struct {
	ID 		int64 `json:"id"`
	RestoId int64 `json:"restoId"`
	Name 	string `json:"name"`
	ImgUrl 	string `json:"imgURL"`
	Status 	int `json:"status"`
}
