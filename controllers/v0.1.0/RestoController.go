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

type RestoController struct {

}

func (controller *RestoController) Save (ctx *gin.Context) {
	fmt.Println(">>> RestoControoler - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.RestoRequesDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	byteReq,_ := json.Marshal(req)
	log.Println("req--> ", string(byteReq))

	res = services.InitializeRestoServiceInterface().Save(&req)

	ctx.JSON(http.StatusOK, res)


}

//func (controller *RestoController) UploadImage (ctx *gin.Context) {
//	fmt.Println(">>> RestoControoler - Upload Image <<<")
//	parent := context.Background()
//	defer parent.Done()
//
//	req := dto.UploadImageReqDto{}
//	res := models.Response{}
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		fmt.Println("Request body error:", err)
//		res.Rc = constants.ERR_CODE_03
//		res.Msg = constants.ERR_CODE_03_MSG
//		ctx.JSON(http.StatusBadRequest, res)
//		return
//	}
//	byteReq,_ := json.Marshal(req)
//	log.Println("req--> ", string(byteReq))
//
//	res = services.InitializeRestoServiceInterface().UploadImage(req)
//
//	ctx.JSON(http.StatusOK, res)
//
//
//}

func (controller *RestoController) GetAll (ctx *gin.Context) {
	fmt.Println(">>> RestoControoler - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	res = services.InitializeRestoServiceInterface().GetAll()

	ctx.JSON(http.StatusOK, res)


}

func (controller *RestoController) GetById (ctx *gin.Context) {
	fmt.Println(">>> RestoControoler - Get By Id <<<")
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


	res = services.InitializeRestoServiceInterface().GetById(int64(id))

	ctx.JSON(http.StatusOK, res)


}

func (controller *RestoController) GetByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> RestoControoler - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.RestoRequesDto{}
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

	res = services.InitializeRestoServiceInterface().GetDataByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)

}

func (controller *RestoController)CheckCodeResto(ctx *gin.Context)  {
	fmt.Println(">>> RestoControoler - CheckCodeResto <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.RestoRequesDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	byteReq,_ := json.Marshal(req)
	log.Println("req--> ", string(byteReq))

	res = services.InitializeRestoServiceInterface().CheckCode(req)

	ctx.JSON(http.StatusOK, res)
}