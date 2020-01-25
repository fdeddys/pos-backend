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

type EMenuItemController struct {

}

func (controller *EMenuItemController) Save (ctx *gin.Context) {

	fmt.Println(">>> EMenuItemController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.MenuItemRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqByte,_ :=json.Marshal(req)
	log.Println("req -> ", string(reqByte))

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

func (controller *EMenuItemController) GetByRestoId(ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By GetByRestoId <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	id, errId := strconv.Atoi(ctx.Param("restoId"))
	if errId != nil {
		log.Println("error", errId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	res = services.InitializeMenuItemServiceInterface().GetByRestoId(int64(id))

	ctx.JSON(http.StatusOK, res)
}

func (controller *EMenuItemController) GetByMenuGroupId (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By GetByMenuGroupId <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	id, errId := strconv.Atoi(ctx.Param("groupId"))
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


func (controller *EMenuItemController) GetByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By GetByFilterPaging <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.MenuItemRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
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

	res = services.InitializeMenuItemServiceInterface().GetDataByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)
}

func (controller *EMenuItemController) UploadImage(ctx *gin.Context)  {
	fmt.Println(">>> EMenuItemController - UploadImage <<<")
	parent := context.Background()
	defer parent.Done()

	var req dto.UploadImageMenuItemRequestDto
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

	res = services.InitializeMenuItemServiceInterface().UploadImage(req)

	ctx.JSON(http.StatusOK, res)

}

func (controller *EMenuItemController) RemoveImage(ctx *gin.Context)  {
	fmt.Println(">>> EMenuItemController - RemoveImage <<<")
	parent := context.Background()
	defer parent.Done()

	var req dto.RemoveImageRequestDto
	var res models.Response

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("reqdata ==>", string(reqByte))

	res = services.InitializeMenuItemServiceInterface().RemoveImage(req)

	ctx.JSON(http.StatusOK, res)

}


/*
func (controller *EMenuItemController) GetByMenuGroupIdAndIdResto (ctx *gin.Context) {
	fmt.Println(">>> EMenuItemController - Get By GetByMenuGroupIdAndIdResto <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	groupId, errGroupId := strconv.Atoi(ctx.Param("groupId"))
	if errGroupId != nil {
		log.Println("error", errGroupId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	restoId, errRestoId := strconv.Atoi(ctx.Param("restoId"))
	if errRestoId != nil {
		log.Println("error", errRestoId)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	log.Println(groupId, restoId)

	res = services.InitializeMenuItemServiceInterface().GetByMenuGroupIdAndRestoId(int64(groupId), int64(restoId))

	ctx.JSON(http.StatusOK, res)
}
*/