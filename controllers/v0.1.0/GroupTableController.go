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

type GroupTableController struct {

}

func (controller *GroupTableController) Save(ctx *gin.Context)  {
	fmt.Println(">>> GroupTableController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.GroupTableRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitGroupTableService().Save(&req)


	ctx.JSON(http.StatusOK, res)
}


func (controller *GroupTableController) GetByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> GroupTableController - Filter Paging <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.GroupTableRequestDto{}
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

	res = services.InitGroupTableService().GetDataByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)


}

func (controller *GroupTableController) Filter (ctx *gin.Context) {
	fmt.Println(">>> GroupTableController - Filter <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.GroupTableRequestDto{}
	res := models.Response{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitGroupTableService().Filter(req)

	ctx.JSON(http.StatusOK, res)


}