package services

import (
	"log"
	"resto-be/constants"
	dbmodels "resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/utils"
)

type UserServiceInterface struct {
}

func InitializeUserServiceInterface() *UserServiceInterface {
	return &UserServiceInterface{}
}

func (service *UserServiceInterface) SaveDataUser(data *dbmodels.User) models.Response {
	var res models.Response

	dataUser := dbmodels.User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		PhoneNumb: data.PhoneNumb,
		Fb:        data.PhoneNumb,
		RestoId: data.RestoId,
	}

	newPass := ""
	if data.ID == 0 {
		newPass = utils.GenerateRandomChar()
		encrPass := utils.HashPassword(data.Email + newPass)
		dataUser.Password = encrPass

	}

	err := repository.SaveUser(&dataUser)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	log.Println("save : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = newPass

	return res

}

// GetDataUserByFilterPaging ...
func (service *UserServiceInterface) GetDataUserByFilterPaging(req dto.UserRequesDto, page int, count int) models.Response {
	var res models.Response

	users, total, err := repository.GetUserFilterPaging(req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = users
	res.TotalData = total

	return res

}
