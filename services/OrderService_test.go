package services

import (
	"log"
	"testing"
)

func TestOrderServiceInterface_GetStatusOrder(t *testing.T) {
	res := InitializeOrderServiceInterface().GetStatusOrder("10")

	log.Println("res --> ", res)
}
