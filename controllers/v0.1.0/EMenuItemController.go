package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
	"strconv"
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

func (controller *EMenuItemController) GetAll (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	res = services.InitializeMenuItemServiceInterface().GetAll()

	ctx.JSON(http.StatusOK, res)


}

func (controller *EMenuItemController) GetById (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By Id <<<")
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
	res = services.InitializeMenuItemServiceInterface().GetById(int64(id))

	ctx.JSON(http.StatusOK, res)
}

func (controller *EMenuItemController) GetByMenuGroupId (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By GetByMenuGroupId <<<")
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
	res = services.InitializeMenuItemServiceInterface().GetByMenuGroupId(int64(id))

	ctx.JSON(http.StatusOK, res)
}