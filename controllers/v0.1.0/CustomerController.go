package v0_1_0

import (
	"context"
	"fmt"
	"net/http"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
}

func (controller *CustomerController) Save(ctx *gin.Context) {
	fmt.Println(">>> CustomerController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dbmodels.Customer{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeCustomerServiceInterface().SaveDataCustomer(&req)

	ctx.JSON(http.StatusOK, res)

}

func (controller *CustomerController) Login(ctx *gin.Context) {
	fmt.Println(">>> Login Customer - Controller <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.LoginRequestDto{}
	res := dto.LoginResponseDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeAuthServiceInterface().AuthLoginCustomer(&req)

	ctx.JSON(http.StatusOK, res)

}
