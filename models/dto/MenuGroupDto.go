package dto

type MenuGroupRequestDto struct {
	ID 		int64 `json:"id"`
	RestoId int64 `json:"restoId"`
	Name 	string `json:"name"`
	ImgUrl 	string `json:"imgURL"`
	JamBuka string `json:"jamBuka"`
	Status 	int `json:"status"`
	RestoCode string `json:"restoCode"`
}


type UploadImageMenuGroupRequestDto struct {
	Data string `json:"data"`
}