package dbmodels

import "encoding/json"

// User ...
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	PhoneNumb string `json:"phoneNumb"`
	Fb        string `json:"fb"`
}

// TableName ...
func (u *User) TableName() string {
	return "public.user"
}

// MarshalJSON ...
func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	usr := user(u)
	usr.Password = "***"
	return json.Marshal(usr)
}
