package repository

import (
	"log"
	"testing"
)

func TestGetOrderDetailByID(t *testing.T)  {
	res := GetOrderDetailByID(9999)
	log.Println("res => ", res)
}
