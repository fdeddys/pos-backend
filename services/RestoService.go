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
		Address: restoDto.Address,
		Desc: restoDto.Desc,
		City: restoDto.City,
		Province: restoDto.Province,
	}

	err := repository.SaveResto(&resto)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		return res
	}

	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = resto


	return res

}