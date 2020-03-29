package dto

type RestoRequesDto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	RestoCode string `json:"restoCode"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	Tax 	int `json:"tax"`
	KetPayment string `json:"ketPayment"`
	ServiceCharge 	int `json:"serviceCharge"`
	//Pictures []ImageDto `json:"pictures"`
	Status 	int `json:"status"`
}

type ImageDto struct {
	ImgUrl		string `json:"imgURL"`
}

type UploadImageRestoRequestDto struct {
	Data string `json:"data"`
	RestoId int64 `json:"restoId"`
	Seq 	int `json:"seq"`
}

