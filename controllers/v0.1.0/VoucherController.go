package v0_1_0

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
	"strconv"
)

type VoucherController struct {
	
}

func (controller *VoucherController) Save(ctx *gin.Context)  {
	fmt.Println(">>> Voucher Controller - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.VoucherRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeVoucherServiceInterface().Save(&req)


	ctx.JSON(http.StatusOK, res)
}

func (controller *VoucherController) GetByCodeVoucher(ctx *gin.Context)  {
	fmt.Println(">>> VoucherController Controoler - GetByCodeVoucher <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.VoucherRequestDto{}
	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeVoucherServiceInterface().GetByCode(req)


	ctx.JSON(http.StatusOK, res)

}

func (controller *VoucherController) GetById (ctx *gin.Context) {
	fmt.Println(">>> VoucherController - Get By Id <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		log.Println("error", errId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}


	res = services.InitializeVoucherServiceInterface().GetById(int64(id))

	ctx.JSON(http.StatusOK, res)

}


func (controller *VoucherController) GetByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> VoucherController - Filter Paging <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.VoucherRequestDto{}
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

	res = services.InitializeVoucherServiceInterface().GetDataByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)


}