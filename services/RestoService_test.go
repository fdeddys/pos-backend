package services

import (
	"log"
	"resto-be/models/dto"
	"testing"
	"github.com/rs/xid"
)

func TestRestoServiceInterface_GenerateFileNameImage(t *testing.T) {

	fileName, url := InitializeRestoServiceInterface().GenerateFileNameImage(12,22)
	log.Println(fileName, url)
}

func TestRestoServiceInterface_RemoveImage(t *testing.T) {
	//req := dto.RemoveImageRestoRequestDto{
	//	ImgUrl: "p3",
	//}
	//res := InitializeRestoServiceInterface().RemoveImage(req)
	//log.Println(res)
}

func TestRestoServiceInterface_UploadImage(t *testing.T) {
	req := dto.UploadImageRestoRequestDto{
		RestoId: 40,
		Seq: 1,
		Data: "data:image/jpg;base64,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",
	}
	res:= InitializeRestoServiceInterface().UploadImage(req)
	log.Println("res--> ", res)
}

func TestGenID(t *testing.T) {
	id := xid.New()
	log.Println("id ==> ",id)

}

func TestRestoServiceInterface_GenerateRestoCode(t *testing.T) {
	res := InitializeRestoServiceInterface().GenerateRestoCode()
	log.Println(res)
}