package dto

type GroupTableRequestDto struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	DiscType  string `json:"discType"`
	Status    int32  `json:"status"`
	RestoCode string `json:"restoCode"`
}
