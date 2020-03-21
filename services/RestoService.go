package services

import (
	"fmt"
	"github.com/rs/xid"
	"log"
	"resto-be/constants"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/hosts/menustorage"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/utils"
)

type RestoServiceInterface struct {
	Send func(menustorage.ReqUploadImageModel)(*menustorage.ResUploadImageModel, error)
}

func InitializeRestoServiceInterface()  *RestoServiceInterface {

	return &RestoServiceInterface{
		Send: menustorage.UploadImage,
	}
}

var (
	hostMinio	string
	bucketNameResto	string
)

func init()  {
	hostMinio = utils.GetEnv("STORAGE_MINIO_URLACCESS", "http://156.67.214.228:9000")
	bucketNameResto = "resto"

}

//func (service *RestoServiceInterface) UploadImage(req dto.UploadImageReqDto) models.Response {
//	fmt.Println("<< RestoServiceInterface - Upload Image >>")
//	var res models.Response
//
//
//
//	return res
//}

func (service *RestoServiceInterface) Save (restoDto *dto.RestoRequesDto) models.Response{
	var res models.Response
	var resto dbmodels.Resto

	resto = dbmodels.Resto{
		ID: restoDto.ID,
		Name: restoDto.Name,
		Address: restoDto.Address,
		RestoCode: restoDto.RestoCode,
		Desc: restoDto.Desc,
		City: restoDto.City,
		Province: restoDto.Province,
		Status: constants.RESTO_ACTIVE,
	}

	if resto.ID == 0 {
		resto.RestoCode = service.GenerateRestoCode()
	}

	// save resto
	err := repository.SaveResto(&resto)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}

	////save pic
	//errSavePicture := service.SavePictures(resto.ID, restoDto.Pictures)
	//if errSavePicture != nil{
	//	log.Println("err save pict : ", err)
	//
	//}
	//
	//log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto


	return res

}

func (service *RestoServiceInterface) GenerateRestoCode() string {

	x:= xid.New().Counter()
	log.Println("x--> ", x)

	restoCode := fmt.Sprintf("%v", x)

	return restoCode
}

func (service *RestoServiceInterface)SavePictures(restoId int64, reqPictures []dto.ImageDto) error {

	repository.DeleteImageRestoByRestoId(restoId)

	for i:=0; i< len(reqPictures); i++ {
		var image dbmodels.RestoPicture

		image.ImgUrl = reqPictures[i].ImgUrl
		image.RestoId = restoId
		image.Status = constants.IMAGE_ACTIVE
		err := repository.SaveImageRestoTemp(&image)
		if err != nil {
			log.Println("err save path image to database : ", err)

			return err
		}
	}

	return nil
}


func (service *RestoServiceInterface) GetAll () models.Response{
	var res models.Response

	restorants, err := repository.GetAllResto()
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = restorants

	return res

}

func (service *RestoServiceInterface) GetById (id int64) models.Response{
	var res models.Response

	resto, err := repository.GetRestoById(id)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto

	return res

}

func (service *RestoServiceInterface) GetDataByFilterPaging (req dto.RestoRequesDto, page int, count int) models.Response{
	var res models.Response

	restorants, total, err := repository.GetRestoFilterPaging(req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = restorants
	res.TotalData = total

	return res

}

func (service *RestoServiceInterface) CheckCode(requesDto dto.RestoRequesDto) models.Response {
	// cek kode resto
	var res models.Response
	resto,_ := repository.GetRestoByRestoCode(requesDto.RestoCode)

	if resto.ID != requesDto.ID{
		if resto.ID > 0 {
			if resto.ID == requesDto.ID{

			}
			res.Rc = constants.ERR_CODE_60
			res.Msg = constants.ERR_CODE_60_MSG

			return res
		}
	}
	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res
}

func (service *RestoServiceInterface) UploadImage (req dto.UploadImageRestoRequestDto) models.Response {
	fmt.Println("<< RestoSErvice -- Upload Image >>")
	var res models.Response

	//// check restoId
	//var resto dbmodels.Resto
	//if req.RestoId == 0 {
	//	resto.Status = constants.RESTO_ACTIVE
	//	err := repository.SaveResto(&resto)
	//	if err != nil {
	//		res.Rc = constants.ERR_CODE_10
	//		res.Msg = constants.ERR_CODE_10_MSG
	//		return res
	//	}
	//	req.RestoId = resto.ID
	//}

	fileName, imgUrl := service.GenerateFileNameImage(req.RestoId,req.Seq)

	log.Println(fileName, imgUrl)


	errUploadPathChan := make(chan error)
	errSendToMinioChan := make(chan error)
	go service.AsyncUploadPath(imgUrl, req.RestoId, errUploadPathChan)

	go service.AsyncSendToMinio(fileName, req.Data, errSendToMinioChan)

	errUploadPath := <-errUploadPathChan
	errSendToMinio := <-errSendToMinioChan

	log.Println("errUploadPath ->", errUploadPath)
	log.Println("errSendToMinio ->", errSendToMinio)
	if errSendToMinio != nil {
		res.Rc = constants.ERR_CODE_21
		res.Msg = constants.ERR_CODE_21_MSG
		return res

	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = imgUrl

	return res
}

func (service *RestoServiceInterface) RemoveImage (req dto.RemoveImageRequestDto) models.Response {
	fmt.Println("<< RestoSErvice -- RemoveImage >>")
	var res models.Response

	pict := repository.GetRestoPictureByImgUrl(req.ImgUrl)
	if pict.ID == 0 {
		log.Println("Image not Found ye")
		res.Rc = constants.ERR_CODE_30
		res.Msg = constants.ERR_CODE_30_MSG
		return res
	}
	err:= repository.RemoveRestoPicture(&pict)
	if err != nil {
		res.Rc = constants.ERR_CODE_12
		res.Msg = constants.ERR_CODE_12_MSG
		return res

	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res
}

func (service *RestoServiceInterface) AsyncSendToMinio (fileName string, data string, errChan chan error)  {

	reqUpload := menustorage.ReqUploadImageModel{
		BucketName: bucketNameResto,
		NameFile: fileName,
		Data: data,
		ContentType: "image/jpeg",
	}

	_, err :=service.Send(reqUpload)
	if err!=nil {
		log.Println("gagal upload")
	}

	errChan <- err
	close(errChan)
	return
}

func (service *RestoServiceInterface) AsyncUploadPath (imgUrl string, restoId int64, errChan chan error)  {

	// get pictureResto by img url
	picture := repository.GetRestoPictureByImgUrl(imgUrl)
	if picture.ID > 0 {
		log.Println("Image sudah ada")
		errChan <- nil
		close(errChan)
		return
	}

	log.Println("create image baru")

	picture.ImgUrl = imgUrl
	picture.RestoId = restoId
	picture.Status = constants.IMAGE_ACTIVE

	err := repository.SaveRestoPicture(&picture)
	if err != nil {
		log.Println("err save path image to database : ", err)

		errChan <- err
		close(errChan)
		return
	}

	errChan <- nil
	close(errChan)
	return

}

// generate filename by idResto and seqnumber
func (service *RestoServiceInterface) GenerateFileNameImage (idResto int64, seq int) (string, string) {

	var fileName string
	var imgUrl string

	fileName = fmt.Sprintf("%v__%v.jpeg", idResto, seq)
	imgUrl = fmt.Sprintf("%v/%v/%v",hostMinio,bucketNameResto,fileName)

	return fileName, imgUrl
}

//
func (service *RestoServiceInterface) SaveByResto(req *dto.RestoRequesDto) models.Response {
	var res models.Response
	//var resto dbmodels.Resto

	log.Println("dto.CurrRestoID", dto.CurrRestoID)
	restoId := dto.CurrRestoID

	resto, err := repository.GetRestoById(restoId)
	if err!=nil {
		res.Rc = constants.ERR_CODE_20
		res.Msg = constants.ERR_CODE_20_MSG

		return res
	}

	resto.Name = req.Name
	resto.Desc = req.Desc
	resto.Address = req.Address
	resto.City = req.City
	resto.Province = req.Province
	resto.Tax = req.Tax
	resto.ServiceCharge = req.ServiceCharge


	if err := repository.SaveResto(&resto); err != nil{
		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto

	return res
}
