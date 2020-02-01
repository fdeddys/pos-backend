package v0_1_0

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
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
	res := dto.LoginCustomerResponseDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeAuthServiceInterface().AuthLoginCustomer(&req)

	fmt.Println("isissss ", res)
	ctx.JSON(http.StatusOK, res)

}

// FilterPage ...
func (controller *CustomerController) FilterPage(ctx *gin.Context) {
	fmt.Println(">>> Customer Controoler - filter Page <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.CustomerDto{}
	res := models.Response{}

	page, errPage := strconv.Atoi(ctx.Param("page"))
	if errPage != nil {
		log.Println("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	count, errCount := strconv.Atoi(ctx.Param("count"))
	if errCount != nil {
		logs.Info("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = services.InitializeCustomerServiceInterface().GetDataCustomerByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)

}

// GetByID ...
func (controller *CustomerController) GetByID(ctx *gin.Context) {
	fmt.Println(">>> Customer - Get By Id <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	id, errID := strconv.Atoi(ctx.Param("id"))
	if errID != nil {
		log.Println("error", errID)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	res = services.InitializeCustomerServiceInterface().GetCustByID(int64(id))

	ctx.JSON(http.StatusOK, res)

}
