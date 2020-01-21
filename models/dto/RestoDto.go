package dto

type RestoRequesDto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	RestoCode string `json:"restoCode"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	Pictures []ImageDto `json:"pictures"`
	Status 	int `json:"status"`
}

type ImageDto struct {
	ID 			int64 `json:"id"`
	ImgUrl		string `json:"imgURL"`
}