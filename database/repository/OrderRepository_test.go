package repository

import (
	"encoding/json"
	"log"
	"resto-be/models/dto"
	"testing"
)

func TestGetOrderDetailByID(t *testing.T)  {
	res := GetOrderDetailByID(9999)
	log.Println("res => ", res)
}

func TestGetOrderDetailByRestoID(t *testing.T) {
	req := dto.OrderRequestDto{
		RestoId: 72,
		StartDate: "2020-03-01",
		EndDate: "2020-03-30",
	}
	res,_:= GetOrderDetailReport(req)
	byteRes,_ := json.Marshal(res)
	log.Println("res -->", string(byteRes))
}
