package packmsg

import (
	"resto-be/hosts/menustorage"
)

func PackMsgMinio(fileName string, bucketName string, data string) menustorage.ReqUploadImageModel  {
	return menustorage.ReqUploadImageModel{
		BucketName: bucketName,
		NameFile: fileName,
		Data: data,
		ContentType: "image/jpeg",
	}
}
