package utils

import (
	"log"
	"testing"
)

func TestGenerateFileNameImage(t *testing.T) {
	fileName, imgUrl :=GenerateFileNameImage("resto")
	log.Println(fileName, imgUrl)
}

func TestGenerateRandomChar(t *testing.T) {
	res := GenerateRandomChar()
	log.Println(res)
}
