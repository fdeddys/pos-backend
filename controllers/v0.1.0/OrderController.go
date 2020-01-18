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

type OrderController struct {

}

func (controller *OrderController) Add (ctx *gin.Context) {
	fmt.Println(">>> OrderController - add <<<")
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

	reqByte,_ :=json.Marshal(req)
	log.Println("req -> ", string(reqByte))
	log.Println("user -> ", dto.CurrUserEmail)
	log.Println("userID -> ", dto.CurrUserID)

	res = services.InitializeOrderServiceInterface().Add(&req)
	resByte,_:= json.Marshal(res)
	log.Println("res add order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)


}