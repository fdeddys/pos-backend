package dto

type RestoRequesDto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	RestoCode string `json:"resto_code"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	Status 	int `json:"status"`
}
