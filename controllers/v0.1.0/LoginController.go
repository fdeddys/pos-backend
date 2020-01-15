package v0_1_0

import (
	"context"
	"fmt"
	"net/http"
	"resto-be/models/dto"
	"resto-be/services"

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
