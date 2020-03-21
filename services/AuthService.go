package services

import (
	"fmt"
	"resto-be/constants"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/models/dto"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//  ..
type AuthServiceInterface struct {
}

// InitializeAuthServiceInterface ..
func InitializeAuthServiceInterface() *AuthServiceInterface {
	return &AuthServiceInterface{}
}

// AuthLogin ...
func (service *AuthServiceInterface) AuthLogin(userDto *dto.LoginRequestDto) dto.LoginResponseDto {
	var res dto.LoginResponseDto

	valReq := service.ValidationRequest(userDto)
	if valReq.Rc != "" {
		return valReq
	}

	user, err := repository.GetUserByEmail(userDto.Email)
	if err != nil {
		res.Rc = constants.ERR_CODE_51
		res.Msg = constants.ERR_CODE_51_MSG
		return res
	}

	valRes := service.ValidationResponse(user, userDto)
	if valRes.Rc != "" {
		return valRes
	}

	token, err := generateToken(user.Email, user.ID, user.RestoId)

	if err != nil {
		res.Rc = constants.ERR_CODE_52
		res.Msg = constants.ERR_CODE_52_MSG
		res.Token = ""
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Token = token
	res.Data = user
	return res
}

// AuthLoginCustomer ...
func (service *AuthServiceInterface) AuthLoginCustomer(userDto *dto.LoginRequestDto) dto.LoginCustomerResponseDto {
	var res dto.LoginCustomerResponseDto

	valReq := service.ValidationRequest(userDto)
	if valReq.Rc != "" {
		res.Rc = valReq.Rc
		res.Msg = valReq.Msg
		return res
	}

	customer, err := repository.GetCustomerByEmail(userDto.Email)
	if err != nil {
		res.Rc = constants.ERR_CODE_51
		res.Msg = constants.ERR_CODE_51_MSG + " " + err.Error()
		return res
	}

	valRes := service.ValidationResponseCustomer(customer, userDto)
	if valRes.Rc != "" {
		fmt.Println("val res ### ", valRes)
		res.Rc = valReq.Rc
		res.Msg = valReq.Msg
		fmt.Println("val res ### ", res)
		return res
	}
	fmt.Println("val res", valRes)
	token, err := generateToken(customer.Email, customer.ID, 0)

	if err != nil {
		res.Rc = constants.ERR_CODE_52
		res.Msg = constants.ERR_CODE_52_MSG
		res.Token = ""
		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Name = customer.Name
	res.Phone = customer.PhoneNumb
	res.CustomerID = fmt.Sprintf("%v", customer.ID)
	res.Msg = constants.ERR_CODE_00_MSG
	res.Token = token
	fmt.Println("finihs")
	return res
}

func generateToken(userEmail string, userId int64, restoId int64) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["userEmail"] = userEmail
	claims["userId"] = fmt.Sprintf("%v", (userId))
	claims["restoId"] = fmt.Sprintf("%v", (restoId))

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000
	timeToString := fmt.Sprintf("%v", umillisec)
	fmt.Println("token Created ", timeToString)
	claims["tokenCreated"] = timeToString

	return sign.SignedString([]byte(constants.TokenSecretKey))

}

func (service *AuthServiceInterface) ValidationRequest(userDto *dto.LoginRequestDto) dto.LoginResponseDto {
	var res dto.LoginResponseDto

	if userDto.Email == "" {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	if userDto.Password == "" {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	return res
}

func (service *AuthServiceInterface) ValidationResponse(user dbmodels.User, userDto *dto.LoginRequestDto) dto.LoginResponseDto {
	var res dto.LoginResponseDto

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

	return res
}

// ValidationResponseCustomer ...
func (service *AuthServiceInterface) ValidationResponseCustomer(customer dbmodels.Customer, userDto *dto.LoginRequestDto) dto.LoginResponseDto {
	var res dto.LoginResponseDto

	if customer.ID == 0 {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	fmt.Println("check pass")
	if customer.Password != userDto.Password {
		res.Rc = constants.ERR_CODE_50
		res.Msg = constants.ERR_CODE_50_MSG
		return res
	}

	return res
}
