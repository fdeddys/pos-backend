package dbmodels

import "encoding/json"

// Customer ...
type Customer struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	PhoneNumb      string `json:"phoneNumb"`
	Fb             string `json:"fb"`
	Status         int    `json:"status"`
	ManualCustomer int    `json:"manual_customer"`
	ManualRestoID  int64  `json:"manual_resto_id"`
}

// TableName ...
func (c *Customer) TableName() string {
	return "public.customer"
}

// MarshalJSON ...
func (c Customer) MarshalJSON() ([]byte, error) {
	type customer Customer // prevent recursion
	cust := customer(c)
	cust.Password = "***"
	return json.Marshal(cust)
}
