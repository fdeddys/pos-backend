package dto

type RestoRequesDto struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	Address string `json:"address"`
}
