package dto

type RoleAccess struct {
	RoleName string `json:"role_name"`
	UserName string `json:"user_name"`
	AccessName string `json:"access_name"`
}
