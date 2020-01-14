package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
)

type MenuGroupInterface struct {

}

func InitializeMenuGroupInterface()  *MenuGroupInterface {
	return &MenuGroupInterface{
	}
}

func (service *MenuGroupInterface) Save (restoDto *dto.MenuGroupRequestDto) models.Response{
	var res models.Response

	menuGroup := dbmodels.MenuGroup{
		ID: restoDto.ID,
		Name: restoDto.Name,
		ImgUrl: restoDto.ImgUrl,
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


func (service *MenuGroupInterface) GetAll () models.Response{
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

func (service *MenuGroupInterface) GetById (id int64) models.Response{
	var res models.Response

	resto, err := repository.GetMenuGroupById(id)
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


func (service *MenuGroupInterface) GetDataByFilterPaging (req dto.MenuGroupRequestDto, page int, count int) models.Response{
	var res models.Response

	restorants, total, err := repository.GetMenuGroupFilterPaging(req, page, count)
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
