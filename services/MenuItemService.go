package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"strconv"
)

type MenuItemServiceInterface struct {
}


func InitializeMenuItemServiceInterface()  *MenuItemServiceInterface {
	return &MenuItemServiceInterface{
	}
}

func (service *MenuItemServiceInterface) Save (reqDto *dto.MenuItemDto) models.Response {
	var res models.Response

	/* BEGIN VALIDATE GROUPID */
	_, errMenuGroup := repository.GetMenuGroupById(reqDto.GroupID)
	if errMenuGroup != nil {
		log.Println("err get from database : ", errMenuGroup)

		res.Rc = constants.ERR_CODE_30
		res.Msg = "MenuGroupID " + strconv.Itoa(int(reqDto.GroupID)) + " "+ constants.ERR_CODE_30_MSG
		return res
	}
	/* END VALIDATE GROUPID */

	/*BEGIN VALIDATE RESTO ID*/
	_, errResto := repository.GetRestoById(reqDto.RestoID)
	if errResto != nil {
		log.Println("err get from database : ", errResto)

		res.Rc = constants.ERR_CODE_30
		res.Msg = "RestoID " + strconv.Itoa(int(reqDto.RestoID)) + " "+ constants.ERR_CODE_30_MSG
		return res
	}
	/*END VALIDATE RESTO ID*/


	menuItem := dbmodels.MenuItem{
		ID: reqDto.ID,
		Name: reqDto.Name,
		ImgUrl: reqDto.ImgUrl,
		Desc: reqDto.Desc,
		GroupID: reqDto.GroupID,
		Price: reqDto.Price,
		RestoID:reqDto.RestoID,
	}


	err := repository.SaveMenuItem(&menuItem)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuItem


	return res


}


func (service *MenuItemServiceInterface) GetAll () models.Response{
	var res models.Response

	menuItems, err := repository.GetAllMenuItem()
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuItems

	return res

}

func (service *MenuItemServiceInterface) GetById (id int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuItemById(id)
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


func (service *MenuItemServiceInterface) GetByMenuGroupId (id int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuItemByMenuGroupId(id)
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

func (service *MenuItemServiceInterface) GetByMenuGroupIdAndRestoId (groupId int64, restoId int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuItemByMenuGroupIdAndRestoId(groupId, restoId)
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
