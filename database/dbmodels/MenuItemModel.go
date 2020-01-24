package dbmodels

type MenuItem struct {
	ID      int64   `json:"id"`
	GroupID int64   `json:"groupId"`
	RestoID int64 	`json:"restoId"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	ImgUrl  string  `json:"imgURL"`
	Price   float64 `json:"price"`
	Stock   int `json:"stock"`
	Status  int     `json:"status"`
}

// TableName ...
func (t *MenuItem) TableName() string {
	return "public.e_menu_item"
}


type MenuItem2 struct {
	ID      int64   `json:"id"`
	GroupID int64   `json:"groupId"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	ImgUrl  string  `json:"imgURL"`
	Price   float64 `json:"price"`
	Stock   int `json:"stock"`
	Status  int     `json:"status"`
}