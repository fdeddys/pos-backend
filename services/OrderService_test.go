package services

import (
	"log"
	"testing"
)

func TestOrderServiceInterface_GetStatusOrder(t *testing.T) {
	res := InitializeOrderServiceInterface().GetStatusOrder("10")

	log.Println("res --> ", res)
}

func TestGenerateOrderNumber(t *testing.T)  {
	res := InitializeOrderServiceInterface().GenerateOrderNumber(78)
	log.Println("res -->", res)

}

func TestReCalculate(t *testing.T) {
	reCalculate(116)
}