package menustorage

type ReqUploadImageModel struct {
	BucketName string
	Data string
	NameFile string
	ContentType string
}

type ResUploadImageModel struct {
	Url string
	NameFile string
	Rc string
	Message string
}