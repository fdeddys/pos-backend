package v0_1_0

import (
	"context"
	"fmt"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"

	"github.com/gin-gonic/gin"
)

type PaymentTypeController struct {
}

// GetAll ...
func (controller *PaymentTypeController) GetAll(ctx *gin.Context) {
	fmt.Println(">>> PaymentTypeController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.PaymentTypeRequestDto{}
	res := models.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitPaymentTypeService().GetAllPaymentType(req)

	ctx.JSON(http.StatusOK, res)

}
