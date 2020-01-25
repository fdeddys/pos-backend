package utils

import (
	"fmt"
	"github.com/rs/xid"
)

var (
	hostMinio	string
)

func init()  {
	hostMinio = GetEnv("STORAGE_MINIO_URLACCESS", "http://156.67.214.228:9000")

}

func GenerateFileNameImage(bucketName string) (string, string) {

	uniqName := xid.New().String()
	fileName := fmt.Sprintf("%v.jpeg", uniqName)

	imgUrl := fmt.Sprintf("%v/%v/%v",hostMinio,bucketName,fileName)

	return fileName, imgUrl
}
