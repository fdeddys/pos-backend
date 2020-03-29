package services

import (
	"fmt"
	"github.com/rs/xid"
	"log"
	"resto-be/constants"
	"resto-be/hosts/menustorage"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
)

type MenuGroupServiceInterface struct {
	Send func(menustorage.ReqUploadImageModel)(*menustorage.ResUploadImageModel, error)

}

func InitializeMenuGroupServiceInterface()  *MenuGroupServiceInterface {
	return &MenuGroupServiceInterface{
		Send: menustorage.UploadImage,
	}
}

func (service *MenuGroupServiceInterface) UploadImage (req dto.UploadImageMenuGroupRequestDto) models.Response {
	fmt.Println("<< MenuGroupServiceInterface -- Upload Image >>")
	var res models.Response

	id := xid.New()

	fileName := fmt.Sprintf("%v.jpeg", id)
	imgUrl := fmt.Sprintf("%v/%v/%v",hostMinio,bucketNameResto,fileName)


	errSendToMinioChan := make(chan error)
	go service.AsyncSendToMinio(fileName, req.Data, errSendToMinioChan)

	errSendToMinio := <-errSendToMinioChan

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

func (service *MenuGroupServiceInterface) Save (menuGroupDto *dto.MenuGroupRequestDto) models.Response{
	var res models.Response

	menuGroup := dbmodels.MenuGroup{
		ID: menuGroupDto.ID,
		Name: menuGroupDto.Name,
		ImgUrl: menuGroupDto.ImgUrl,
		JamBuka: menuGroupDto.JamBuka,
		RestoId: menuGroupDto.RestoId,
		Status: menuGroupDto.Status,
	}

	err := repository.SaveMenuGroup(&menuGroup)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuGroup


	log.Println(menuGroup)
	return res
}


func (service *MenuGroupServiceInterface) GetAll () models.Response{
	var res models.Response

	restorants, err := repository.GetAllMenuGroup()
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

func (service *MenuGroupServiceInterface) GetById (id int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuGroupById(id)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuGroup

	return res

}


func (service *MenuGroupServiceInterface) GetDataByFilterPaging (req dto.MenuGroupRequestDto, restoId int64, page int, count int) models.Response{
	var res models.Response

	menuGroups, total, err := repository.GetMenuGroupFilterPaging(req, restoId, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuGroups
	res.TotalData = total

	return res

}


func (service *MenuGroupServiceInterface) GetByIdResto (GetByIdResto int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuGroupByIdResto(GetByIdResto)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuGroup

	return res

}

func (service *MenuGroupServiceInterface) AsyncSendToMinio (fileName string, data string, errChan chan error)  {

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