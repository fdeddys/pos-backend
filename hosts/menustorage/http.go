package menustorage

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"resto-be/utils"
)

var (
	host					string
	endpointUpload			string
)

func init()  {
	host = utils.GetEnv("STORAGE_MINIO_URLACCESS", "http://156.67.214.228:9000")
	endpointUpload = utils.GetEnv("STORAGE_MINIO_UPLOAD", "/upload")
}

func UploadImage(req ReqUploadImageModel) (*ResUploadImageModel, error)  {
	var resp ResUploadImageModel

	url := host + endpointUpload

	data, err := utils.HTTPPost(url, req)
	if err != nil {
		logs.Error("upload image ", err.Error())
		return &resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		logs.Error("Failed to unmarshaling response upload from minioapp ", err.Error())
		return &resp, err
	}

	return &resp, err
}