package dto

// AddOrderItemDto ...
// add order item to Tabel
type AddOrderItemDto struct {
	TableID int64 `json:"tableId"`
	ItemID  int64 `json:"itemId"`
}
