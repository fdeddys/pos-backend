package dto

type OrderRequestDto struct {
	ID 		int64 `json:"id"`
	OrderNo	string `json:"orderNo"`
	TableId int64 `json:"tableId"`
	RestoId int64 `json:"restoId"`
	OrderDetails []OrderDetailRequest `json:"orderDetails"`

}

type OrderDetailRequest struct {
	EMenuItem int64 `json:"eMenuItem"`
	Qty int `json:"qty"`
}