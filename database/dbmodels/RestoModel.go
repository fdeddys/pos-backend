package dbmodels

type Resto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	RestoCode string `json:"restoCode"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	Status 	int `json:"status"`
	RestoPicture []RestoPicture `json:"restoPicture"gorm:"foreignkey:RestoId"`
}


// TableName ...
func (t *Resto) TableName() string {
	return "public.resto"
}
