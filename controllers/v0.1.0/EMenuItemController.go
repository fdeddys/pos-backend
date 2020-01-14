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

type EMenuItemController struct {

}

func (controller *EMenuItemController) Save (ctx *gin.Context) {

	fmt.Println(">>> EMenuItemController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.MenuItemDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeMenuItemServiceInterface().Save(&req)


	ctx.JSON(http.StatusOK, res)

}