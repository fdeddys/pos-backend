package v0_1_0

import (
	"context"
	"encoding/json"
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

type EMenuGroupController struct {

}

func (controller *EMenuGroupController) UploadImage(ctx *gin.Context)  {
	fmt.Println(">>> EMenuGroupController - UploadImage <<<")
	parent := context.Background()
	defer parent.Done()

	var req dto.UploadImageMenuGroupRequestDto
	var res models.Response

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	req2:= req
	req2.Data = "--gambar base64--"
	reqByte,_ := json.Marshal(req2)
	log.Println("reqdata ==>", string(reqByte))

	res = services.InitializeMenuGroupServiceInterface().UploadImage(req)

	ctx.JSON(http.StatusOK, res)

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

	res = services.InitializeMenuGroupServiceInterface().Save(&req)


	ctx.JSON(http.StatusOK, res)

}

func (controller *EMenuGroupController) GetAll (ctx *gin.Context) {
	fmt.Println(">>> EMenuGroupController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	res = services.InitializeMenuGroupServiceInterface().GetAll()

	ctx.JSON(http.StatusOK, res)


}


func (controller *EMenuGroupController) GetById (ctx *gin.Context) {
	fmt.Println(">>> EMenuGroupController - Get By Id <<<")
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


	res = services.InitializeMenuGroupServiceInterface().GetById(int64(id))

	ctx.JSON(http.StatusOK, res)


}


func (controller *EMenuGroupController) GetByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> EMenuGroupController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.MenuGroupRequestDto{}
	res := models.Response{}

	restoId, errRestoId := strconv.Atoi(ctx.Param("restoId"))
	if errRestoId != nil {
		log.Println("error", errRestoId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

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

	res = services.InitializeMenuGroupServiceInterface().GetDataByFilterPaging(req, int64(restoId), page, count)

	ctx.JSON(http.StatusOK, res)


}


func (controller *EMenuGroupController) GetByIdResto (ctx *gin.Context) {
	fmt.Println(">>> EMenuGroupController - GetByIdResto <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	restoId, errRestoId := strconv.Atoi(ctx.Param("restoId"))
	if errRestoId != nil {
		log.Println("error", errRestoId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}


	res = services.InitializeMenuGroupServiceInterface().GetByIdResto(int64(restoId))

	ctx.JSON(http.StatusOK, res)


}