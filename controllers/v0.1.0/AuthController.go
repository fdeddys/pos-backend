package v0_1_0

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/models/dto"
	"resto-be/models"
	"resto-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

// Login ...
func (controller *AuthController) Login(ctx *gin.Context) {
	fmt.Println(">>> Login - Controller <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.LoginRequestDto{}
	res := dto.LoginResponseDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeAuthServiceInterface().AuthLogin(&req)

	ctx.JSON(http.StatusOK, res)

}


// Reset Password
func (controller *AuthController) ResetPassword (ctx *gin.Context) {
	fmt.Println(">>> AuthControoler - Reset Password by User id<<<")
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

	restoId := dto.CurrRestoID
	if restoId != 0 {
		res.Rc = constants.ERR_CODE_20
		res.Msg = constants.ERR_CODE_20_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}


	res = services.InitializeAuthServiceInterface().ResetPasswordByIdUser(int64(id))

	ctx.JSON(http.StatusOK, res)


}
