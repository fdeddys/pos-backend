package dto

type OrderRequestDto struct {
	ID           int64                `json:"id"`
	OrderNo      string               `json:"orderNo"`
	TableId      int64                `json:"tableId"`
	RestoId      int64                `json:"restoId"`
	RestoCode    string                `json:"restoCode"`
	CustomerId   int64                `json:"customerId"`
	OrderDetails []OrderDetailRequest `json:"orderDetails"`
	Total        int64                `json:"total"`
	Notes        string               `json:"notes"`
	Status       string               `json:"status"`
	StartDate 	 string				  `json:"startDate"`
	EndDate		 string 			  `json:"endDate"`
	PaymentStatus string 			  `json:"paymentStatus"`
}

type OrderDetailRequest struct {
	EMenuItem int64 `json:"eMenuItem"`
	Qty       int   `json:"qty"`
	ID        int64 `json:"id"`
}
