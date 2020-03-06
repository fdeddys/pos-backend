package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-be/models"
	"resto-be/services"
)

type CategoryController struct {

}

func (controller *CategoryController) GetAll(ctx *gin.Context)  {
	fmt.Println(">>> CategoryController - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	res = services.InitCategoryServiceInterface().GetAll()

	ctx.JSON(http.StatusOK, res)
}