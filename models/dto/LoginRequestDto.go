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
}
