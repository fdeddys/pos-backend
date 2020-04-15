package dto

type TableRequestDto struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int32  `json:"status"`
	GroupTabelID string `json:"groupTabelId"`
	RestoCode string `json:"restoCode"`
}