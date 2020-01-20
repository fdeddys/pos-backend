package v0_1_0

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
)

type ImageController struct {
}

func (controller *ImageController) Upload(ctx *gin.Context) {
	fmt.Println(">>> User Controoler - Get All <<<")
	parent := context.Background()
	defer parent.Done()
	req := dto.UploadImageReqDto{}
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

	res = services.InitializeImageServiceInterface().Upload(req)

	ctx.JSON(http.StatusOK, res)
}