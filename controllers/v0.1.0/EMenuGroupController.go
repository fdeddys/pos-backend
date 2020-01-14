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

type EMenuGroupController struct {

}

func (controller *EMenuGroupController) Save (ctx *gin.Context) {
	fmt.Println(">>> EMenuGroupController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.MenuGroupRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeMenuGroupInterface().Save(&req)


	ctx.JSON(http.StatusOK, res)

}