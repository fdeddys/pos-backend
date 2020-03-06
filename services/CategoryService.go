package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
)

type CategoryServiceInterface struct {

}

func InitCategoryServiceInterface() *CategoryServiceInterface {
	return &CategoryServiceInterface{}
}

func (service *CategoryServiceInterface) GetAll () models.Response {
	var res models.Response

	categories, err := repository.GetAllCategory()

	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = categories

	return res
}