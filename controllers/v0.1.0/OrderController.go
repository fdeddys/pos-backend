package v0_1_0

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/report"
	"resto-be/services"

	"github.com/astaxie/beego/logs"
)

type OrderController struct {
}

func (controller *OrderController) GetByCustPage(ctx *gin.Context) {
	fmt.Println(">>> Order Controoler - Get by cust PAGE <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.OrderRequestDto{}
	res := models.Response{}

	page, errPage := strconv.Atoi(ctx.Param("page"))
	if errPage != nil {
		log.Println("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	count, errCount := strconv.Atoi(ctx.Param("count"))
	if errCount != nil {
		logs.Info("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)
	res = services.InitializeOrderServiceInterface().GetByCustomerPage(&req, page, count)

	ctx.JSON(http.StatusOK, res)

}

func (controller *OrderController) Add(ctx *gin.Context) {
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

	reqByte, _ := json.Marshal(req)
	log.Println("req -> ", string(reqByte))
	log.Println("user -> ", dto.CurrUserEmail)
	log.Println("userID -> ", dto.CurrUserID)

	res = services.InitializeOrderServiceInterface().Add(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res add order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}

func (controller *OrderController) PrintPreview(c *gin.Context) {

	orderID, errPage := strconv.ParseInt(c.Param("id"), 10, 64)
	if errPage != nil {
		logs.Info("error", errPage)
		c.JSON(http.StatusBadRequest, "id not supplied")
		c.Abort()
		return
	}

	// fmt.Println("-------->", req)

	report.GenerateReceiveReport(orderID)

	header := c.Writer.Header()
	header["Content-type"] = []string{"application/x-pdf"}
	header["Content-Disposition"] = []string{"attachment; filename= invoice.pdf"}

	file, _ := os.Open("invoice.pdf")

	io.Copy(c.Writer, file)
	return
}

// GetById ...
func (controller *OrderController) GetByID(ctx *gin.Context) {
	fmt.Println(">>> Order Controoler - Get by ID <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "id not supplied")
		ctx.Abort()
		return
	}

	res = services.InitializeOrderServiceInterface().GetById(orderID)

	ctx.JSON(http.StatusOK, res)

}

// GetOrderDetail ...
func (controller *OrderController) GetOrderDetail(ctx *gin.Context) {
	fmt.Println(">>> Order Controoler - Get order detail by Order ID <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	orderID, err := strconv.ParseInt(ctx.Param("orderId"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "id not supplied")
		ctx.Abort()
		return
	}

	res = services.InitializeOrderServiceInterface().GetOrderDetailByOrderID(orderID)

	ctx.JSON(http.StatusOK, res)

}
