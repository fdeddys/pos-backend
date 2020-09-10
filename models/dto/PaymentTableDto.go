package dto

// PaymentByTabel ...
type PaymentByTabel struct {
	TableID int64 `json:"tableId"`
	Debit   int64 `json:"debit"`
	Kredit  int64 `json:"kredit"`
	Cash    int64 `json:"cash"`
	Other   int64 `json:"other"`
}
