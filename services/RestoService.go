package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
)

type RestoServiceInterface struct {

}

func InitializeRestoServiceInterface()  *RestoServiceInterface {
	return &RestoServiceInterface{
	}
}

func (service *RestoServiceInterface) Save (restoDto *dto.RestoRequesDto) models.Response{
	var res models.Response

	resto := dbmodels.Resto{
		ID: restoDto.ID,
		Name: restoDto.Name,
		RestoCode: restoDto.RestoCode,
		Address: restoDto.Address,
		Desc: restoDto.Desc,
		City: restoDto.City,
		Province: restoDto.Province,
		Status: constants.RESTO_ACTIVE,
	}

	// save resto
	err := repository.SaveResto(&resto)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}

	//save pic
	errSavePicture := service.SavePictures(resto.ID, restoDto.Pictures)
	if errSavePicture != nil{
		log.Println("err save pict : ", err)

	}

	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto


	return res

}

func (service *RestoServiceInterface)SavePictures(restoId int64, reqPictures []dto.ImageDto) error {

	repository.DeleteImageRestoByRestoId(restoId)

	for i:=0; i< len(reqPictures); i++ {
		var image dbmodels.RestoPicture

		image.ImgUrl = reqPictures[i].ImgUrl
		image.RestoId = restoId
		image.Status = constants.IMAGE_ACTIVE
		err := repository.SaveImageResto(&image)
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
