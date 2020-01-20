package dto

type UploadImageReqDto struct {
	BucketName string `json:"bucketName"`
	Data string `json:"data"`
	NameFile string `json:"nameFile"`
	ContentType string `json:"contentType"`
}