package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
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





	ctx.JSON(http.StatusOK, res)


}
