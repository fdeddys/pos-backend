package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
)

type VoucherController struct {
	
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

	res = services.InitializeVoucherServiceInterface().GetByCode(req.Code)


	ctx.JSON(http.StatusOK, res)

}