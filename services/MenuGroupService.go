package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
)

type MenuGroupServiceInterface struct {

}

func InitializeMenuGroupServiceInterface()  *MenuGroupServiceInterface {
	return &MenuGroupServiceInterface{
	}
}

func (service *MenuGroupServiceInterface) Save (menuGroupDto *dto.MenuGroupRequestDto) models.Response{
	var res models.Response

	menuGroup := dbmodels.MenuGroup{
		ID: menuGroupDto.ID,
		Name: menuGroupDto.Name,
		ImgUrl: menuGroupDto.ImgUrl,
		RestoId: menuGroupDto.RestoId,
		Status: constants.GROUP_ACTIVE,
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