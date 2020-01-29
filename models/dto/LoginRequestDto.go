package dto

// LoginRequestDto ...
type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponseDto ...
type LoginResponseDto struct {
	Rc    string `json:"rc"`
	Msg   string `json:"message"`
	Token string `json:"token"`
	Data interface{} `json:"data"`
}

// LoginCustomerResponseDto ...
type LoginCustomerResponseDto struct {
	Rc         string `json:"rc"`
	Msg        string `json:"message"`
	CustomerID string `json:"customerId"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Token      string `json:"token"`
}
