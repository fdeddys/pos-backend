package services

import (
	"encoding/json"
	"fmt"
	"log"
	"resto-be/constants"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/hosts/menustorage"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/utils"
	"resto-be/utils/packmsg"
	"strconv"
)

type MenuItemServiceInterface struct {
	Send func(menustorage.ReqUploadImageModel)(*menustorage.ResUploadImageModel, error)

}

func InitializeMenuItemServiceInterface()  *MenuItemServiceInterface {
	return &MenuItemServiceInterface{
		Send: menustorage.UploadImage,
	}
}

func (service *MenuItemServiceInterface) Save (reqDto *dto.MenuItemRequestDto) models.Response {
	var res models.Response



	/* BEGIN VALIDATE GROUPID */
	menuGroup, errMenuGroup := repository.GetMenuGroupById(reqDto.GroupID)
	if errMenuGroup != nil {
		log.Println("err get from database : ", errMenuGroup)

		res.Rc = constants.ERR_CODE_30
		res.Msg = "MenuGroupID " + strconv.Itoa(int(reqDto.GroupID)) + " "+ constants.ERR_CODE_30_MSG
		return res
	}
	/* END VALIDATE GROUPID */


	/*
	* TODO VALIDATE RESTO
	*/
	/* to
	if dto.CurrRestoID != 0 {
		menuGroup, errMenuGroup := repository.GetMenuGroupById(reqDto.GroupID)
		if errMenuGroup != nil {
			log.Println("err get from database : ", errMenuGroup)

			res.Rc = constants.ERR_CODE_30
			res.Msg = "MenuGroupID " + strconv.Itoa(int(reqDto.GroupID)) + " "+ constants.ERR_CODE_30_MSG
			return res
		}

		// validasi user berhak/tidak untuk save by resto_id
		if menuGroup.RestoId != dto.CurrRestoID{
			res.Rc = constants.ERR_CODE_20
			res.Msg = constants.ERR_CODE_20_MSG
			return res
		}
	}
	/* END VALIDATE RESTO */

	menuItem := dbmodels.MenuItem{
		ID: reqDto.ID,
		Name: reqDto.Name,
		ImgUrl: reqDto.ImgUrl,
		Desc: reqDto.Desc,
		GroupID: reqDto.GroupID,
		RestoID: menuGroup.RestoId,
		Price: reqDto.Price,
		Stock:reqDto.Stock,
		IsFavorite: reqDto.IsFavorite,
		CategoryId: reqDto.CategoryId,
		Status: reqDto.Status,
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



	log.Println("get data : ", menuGroup)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuGroup

	return res

}

func (service *MenuItemServiceInterface) GetByRestoId (id int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuItemByRestoId(id)
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

func (service *MenuItemServiceInterface) GetFavoriteByRestoId (id int64) models.Response{
	var res models.Response

	menuGroup, err := repository.GetMenuItemFavoriteByRestoId(id)
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

func (service *MenuItemServiceInterface) GetDataByFilterPaging (req dto.MenuItemRequestDto, page int, count int) models.Response{
	fmt.Println(">>> MenuItemServiceInterface - GetDataByFilterPaging <<<")
	var res models.Response

	reqByte,_ := json.Marshal(req)
	log.Println("reqData -> ", string(reqByte))
	if req.RestoId == 0 {
		res.Rc = constants.ERR_CODE_40
		res.Msg = constants.ERR_CODE_40_MSG
		return res
	}


	menuItems, total, err := repository.GetMenuItemFilterPaging(req, page, count)
	if err != nil {
		log.Println("err get menu items from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = menuItems
	res.TotalData = total

	return res

}

func (service *MenuItemServiceInterface) UploadImage(req dto.UploadImageMenuItemRequestDto) models.Response {
	fmt.Println("<< MenuItemService -- Upload Image >>")
	var res models.Response

	fileName, imgUrl := utils.GenerateFileNameImage(constants.BUCKET_MENU_ITEM)
	log.Println(fileName, imgUrl)

	erPathImageChan := make(chan error)
	errSendToMinioChan := make(chan error)
	go service.AsyncSavePathImage(imgUrl, req.ID, req.MenuItemId, erPathImageChan)
	go service.AsyncSendToMinio(fileName, req.Data, errSendToMinioChan)

	erPathImage := <-erPathImageChan
	errSendToMinio := <-errSendToMinioChan

	log.Println("errUploadPath ->", erPathImage)
	log.Println("errSendToMinio ->", errSendToMinio)
	if erPathImage != nil {
		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
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

func (service *MenuItemServiceInterface) RemoveImage (req dto.RemoveImageRequestDto) models.Response {
	fmt.Println("<< MenuItemServiceInterface -- RemoveImage >>")
	var res models.Response

	pict := repository.GetMenuItemPictureByImgUrl(req.ImgUrl)
	if pict.ID == 0 {
		log.Println("Image not Found ye")
		res.Rc = constants.ERR_CODE_30
		res.Msg = constants.ERR_CODE_30_MSG
		return res
	}

	err:= repository.RemoveMenuItemPicture(&pict)
	if err != nil {
		res.Rc = constants.ERR_CODE_12
		res.Msg = constants.ERR_CODE_12_MSG
		return res

	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res
}

func (service *MenuItemServiceInterface) AsyncSavePathImage(imgUrl string, id int64, menuItemId int64, errChan chan error)  {
	picture := dbmodels.MenuItemPicture{
		ID: id,
		Status:constants.IMAGE_ACTIVE,
		ImgUrl: imgUrl,
		MenuItemId: menuItemId,
	}

	err := repository.SaveMenuItemPicture(&picture)
	if err != nil {
		log.Println("err save path image item to database : ", err)

		errChan <- err
		close(errChan)
		return
	}

	errChan <- nil
	close(errChan)
	return

}

func (service *MenuItemServiceInterface) AsyncSendToMinio (fileName string, data string, errChan chan error)  {

	reqUpload := packmsg.PackMsgMinio(fileName, constants.BUCKET_MENU_ITEM, data)

	_, err :=service.Send(reqUpload)
	if err!=nil {
		log.Println("gagal upload")
	}

	errChan <- err
	close(errChan)
	return
}

func (service *MenuItemServiceInterface) Filter (req dto.MenuItemRequestDto) models.Response{
	fmt.Println(">>> MenuItemServiceInterface - Filter <<<")
	var res models.Response

	reqByte,_ := json.Marshal(req)
	log.Println("reqData -> ", string(reqByte))

	menuItems, err := repository.GetMenuItemFilter(req)
	if err != nil {
		log.Println("err get menu items from database : ", err)

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


/*
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
*/