package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

type TableService struct {
}

func InitTableService() *TableService {
	return &TableService{}
}

func (service *TableService) Save (req *dto.TableRequestDto) models.Response{
	var res models.Response

	table := dbmodels.Table{
		ID: req.ID,
		Name: req.Name,
		Status: req.Status,
		GroupTabelID: req.GroupTabelID,

	}

	err := repository.SaveTable(&table)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = table


	log.Println(table)
	return res
}

func (service *TableService) Filter (req dto.TableRequestDto) models.Response{
	var res models.Response


	tables, err := repository.GetTableFilter(req)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = tables

	return res

}