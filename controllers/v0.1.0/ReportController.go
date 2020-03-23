package v0_1_0

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
	"strconv"
)

type ReportController struct {
}


func (controller *ReportController) Order(ctx *gin.Context)  {
	fmt.Println(">>> Report Controoler - Order <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
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

	res, filePath := services.InitializeReportServiceInterface().Order(&req)
	if res.Rc == constants.ERR_CODE_00 {
		w := ctx.Writer
		f, err := os.Open(filePath)
		if f != nil {
			defer f.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Failed to open file", err)
			return
		}
		contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
		w.Header().Set("Content-Disposition", contentDisposition)

		if _, err := io.Copy(w, f); err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Error(w,"File Not Found", http.StatusInternalServerError)
			fmt.Println("Failed to copy file", err)
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ReportController) OrderDetail(ctx *gin.Context)  {
	fmt.Println(">>> Report Controoler - OrderDetail <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
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

	res, filePath := services.InitializeReportServiceInterface().OrderDetail(&req)
	if res.Rc == constants.ERR_CODE_00 {
		w := ctx.Writer
		f, err := os.Open(filePath)
		if f != nil {
			defer f.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Failed to open file", err)
			return
		}
		contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
		w.Header().Set("Content-Disposition", contentDisposition)

		if _, err := io.Copy(w, f); err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Error(w,"File Not Found", http.StatusInternalServerError)
			fmt.Println("Failed to copy file", err)
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}


func (controller *ReportController) GetOrderByFilterPaging (ctx *gin.Context) {
	fmt.Println(">>> ReportController - Get By GetOrderByFilterPaging <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
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

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)
	req.RestoId = dto.CurrRestoID
	res = services.InitializeOrderServiceInterface().GetByFilterPaging(&req, page, count)

	ctx.JSON(http.StatusOK, res)

}
