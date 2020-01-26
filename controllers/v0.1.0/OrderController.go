package v0_1_0

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"

	"github.com/astaxie/beego/logs"
)

type OrderController struct {
}

func (controller *OrderController) GetByCustPage(ctx *gin.Context) {
	fmt.Println(">>> Order Controoler - Get by cust PAGE <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
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

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)
	res = services.InitializeOrderServiceInterface().GetByCustomerPage(&req, page, count)

	ctx.JSON(http.StatusOK, res)

}

func (controller *OrderController) Add(ctx *gin.Context) {
	fmt.Println(">>> OrderController - add <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqByte, _ := json.Marshal(req)
	log.Println("req -> ", string(reqByte))
	log.Println("user -> ", dto.CurrUserEmail)
	log.Println("userID -> ", dto.CurrUserID)

	res = services.InitializeOrderServiceInterface().Add(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res add order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}
