package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
)

func LoginController(ctx *gin.Context)  {
	fmt.Println(">>> Login - Controller <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.LoginRequestDto{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeUserServiceInterface().AuthLogin(&req)


	ctx.JSON(http.StatusOK, res)

}