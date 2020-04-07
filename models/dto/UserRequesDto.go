package dto

// UserRequesDto ...
type UserRequesDto struct {
	Name string `json:"name"`
}

type ChangePasswordDto struct {
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}
