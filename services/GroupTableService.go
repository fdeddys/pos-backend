package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
)

type GroupTableService struct {
}

func InitGroupTableService() *GroupTableService {
	return &GroupTableService{}
}

func (service *GroupTableService) Save (req *dto.GroupTableRequestDto) models.Response{
	var res models.Response


	resto, errResto := repository.GetRestoById(dto.CurrRestoID)
	if errResto != nil {
		res.Rc = constants.ERR_CODE_20
		res.Msg = constants.ERR_CODE_20_MSG
		return res
	}

	groupTable := dbmodels.GroupTable{
		ID: req.ID,
		Name: req.Name,
		Status: req.Status,
		RestoCode: resto.RestoCode,

	}

	err := repository.SaveGroupTable(&groupTable)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = groupTable


	log.Println(groupTable)
	return res
}

func (service *GroupTableService) GetDataByFilterPaging (req dto.GroupTableRequestDto, page int, count int) models.Response{
	var res models.Response

	resto, errResto := repository.GetRestoById(dto.CurrRestoID)
	if errResto != nil {
		res.Rc = constants.ERR_CODE_20
		res.Msg = constants.ERR_CODE_20_MSG
		return res
	}

	req.RestoCode = resto.RestoCode

	groupTables, total, err := repository.GetGroupTablePaging(req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = groupTables
	res.TotalData = total

	return res

}

func (service *GroupTableService) Filter (req dto.GroupTableRequestDto) models.Response{
	var res models.Response

	groupTables, err := repository.GetGroupTableFilter(req)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = groupTables
	return res

}