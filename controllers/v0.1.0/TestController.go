package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-be/models"
)

func TestController(ctx *gin.Context)  {
	fmt.Println(">>> Test - Controller <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{
		Rc:  "00",
		Msg: "Transaction succesa",
	}
	ctx.JSON(http.StatusOK, res)
}
