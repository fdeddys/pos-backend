package services

import (
	"resto-be/constants"
	"resto-be/hosts/menustorage"
	"resto-be/models"
	"resto-be/models/dto"
)

type ImageServiceInterface struct {
	Send func(menustorage.ReqUploadImageModel)(*menustorage.ResUploadImageModel, error)
}

func InitializeImageServiceInterface()  *ImageServiceInterface {
	return &ImageServiceInterface{
		Send: menustorage.UploadImage,
	}
}

func (service *ImageServiceInterface) Upload (req dto.UploadImageReqDto) models.Response {
	var res models.Response

	reqUpload := menustorage.ReqUploadImageModel{
		BucketName: req.BucketName,
		NameFile: req.NameFile,
		Data: req.Data,
		ContentType: "image/jpeg",
	}

	_, err :=service.Send(reqUpload)
	if err != nil {
		res.Rc = constants.ERR_CODE_21
		res.Msg = constants.ERR_CODE_21_MSG
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res
}