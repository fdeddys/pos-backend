package repository

import (
	"log"
	"testing"
)

func TestGetRestoPictureByImgUrl(t *testing.T) {
	res := GetRestoPictureByImgUrl("tp://156.67.214.228:9000/resto/Xdjdjdj0.jpeg")
	log.Println( res)
}
