package dto

type UploadImageReqDtotemp struct {
	BucketName string `json:"bucketName"`
	Data string `json:"data"`
	NameFile string `json:"nameFile"`
}