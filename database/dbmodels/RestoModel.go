package dbmodels

type Resto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
}


// TableName ...
func (t *Resto) TableName() string {
	return "public.resto"
}
