package services

import (
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
)

//  ..
type UserServiceInterface struct {
}

// InitializeUserServiceInterface ..
func InitializeUserServiceInterface()  *UserServiceInterface {
	return &UserServiceInterface{
	}
}

func (service *UserServiceInterface) AuthLogin(userDto *dto.LoginRequestDto) models.Response  {
	var res models.Response

	if userDto.Username == "" {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	if userDto.Password == "" {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = "token"

	return res
}
