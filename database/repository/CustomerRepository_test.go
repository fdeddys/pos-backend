package repository

import (
	"log"
	"testing"
)

func TestGetCustomerEmailAndRestoId(t *testing.T) {
	x, err:=GetCustomerEmailAndRestoId("aaaa", 722)
	log.Println(x)
	log.Println(err)
}
