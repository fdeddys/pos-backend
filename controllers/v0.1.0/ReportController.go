package v0_1_0

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
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
