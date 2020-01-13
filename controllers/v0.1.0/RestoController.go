package v0_1_0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RestoController(ctx *gin.Context) {
	fmt.Println(">>> Login - Controller <<<")
	parent := context.Background()
	defer parent.Done()

}
