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