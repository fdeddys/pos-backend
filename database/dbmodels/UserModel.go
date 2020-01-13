package dbmodels

type User struct {
	ID 		int64 `json:"id"`
	Name 	string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumb	string `json:"phone_numb"`
	Fb 		string `json:"fb"`
}

// TableName ...
func (t *User) TableName() string {
	return "public.user"
}
