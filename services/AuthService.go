package services

import (
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
)

//  ..
type AuthServiceInterface struct {
}

// InitializeAuthServiceInterface ..
func InitializeAuthServiceInterface()  *AuthServiceInterface {
	return &AuthServiceInterface{
	}
}

func (service *AuthServiceInterface) AuthLogin(userDto *dto.LoginRequestDto) models.Response  {
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

	user, err := repository.GetUserByEmail(userDto.Username)

	if err != nil {
		res.Rc = constants.ERR_CODE_51
		res.Msg = constants.ERR_CODE_51_MSG
		return res
	}
	// fmt.Println("USER ---> ", user)
	if user.ID == 0 {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	if user.Password != userDto.Password {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}


	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = "token"

	return res
}
