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
func (controller *OrderController) GetOrderDetailByOrderId(ctx *gin.Context) {
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

func (controller *OrderController) GetByFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> OrderController - Get By GetByFilterPaging <<<")
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
	res = services.InitializeOrderServiceInterface().GetByFilterPaging(&req, page, count)

	ctx.JSON(http.StatusOK, res)

}

func (controller *OrderController) UpdateStatusOrderDetail(ctx *gin.Context) {
	fmt.Println(">>> OrderController -  UpdateStatusOrderDetail <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}
	req := dto.OrderRequestDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)

	switch req.Status {
	case constants.COOK_COOKING_DESC:
		req.Status = constants.COOK_COOKING
		res = services.InitializeOrderServiceInterface().UpdateCookStatus(&req)
	case constants.COOK_DELIVERY_DESC:
		req.Status = constants.COOK_DELIVERY
		res = services.InitializeOrderServiceInterface().UpdateCookStatus(&req)
	case constants.COOK_AT_LOCATION_DESC:
		req.Status = constants.COOK_AT_LOCATION
		res = services.InitializeOrderServiceInterface().UpdateCookStatus(&req)
	case constants.COOK_ON_HAND_DESC:
		req.Status = constants.COOK_ON_HAND
		res = services.InitializeOrderServiceInterface().UpdateCookStatus(&req)
	case constants.COOK_CANCEL_DESC:
		req.Status = constants.COOK_CANCEL
		res = services.InitializeOrderServiceInterface().UpdateCookStatus(&req)
	}

	// logs.Info(statusPayment)
	// res = services.InitializeOrderServiceInterface().Add(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res update pay order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}

// UpdateStatusPayment ...
func (controller *OrderController) UpdateStatusOrder(ctx *gin.Context) {
	fmt.Println(">>> OrderController - Update status <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}
	req := dto.OrderRequestDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)

	switch req.Status {
	case "PAID":
		req.Status = constants.PAID
		res = services.InitializeOrderServiceInterface().UpdatePayment(&req)
	case "CANCEL":
		req.Status = constants.CANCEL
		res = services.InitializeOrderServiceInterface().UpdatePayment(&req)
	case "cook":
		// statusPayment = constants.PAID
	case "delivery":
		// statusPayment = constants.CANCEL
	}

	// logs.Info(statusPayment)
	// res = services.InitializeOrderServiceInterface().Add(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res update pay order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}

// UpdateStatusCompleteOrder ...
func (controller *OrderController) UpdateStatusCompleteOrder(ctx *gin.Context) {
	fmt.Println(">>> OrderController - UpdateStatusCompleteOrder <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}
	req := dto.OrderRequestDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)

	switch req.Status {
	case constants.COMPLETE_DESC:
		req.Status = constants.COMPLETE
		res = services.InitializeOrderServiceInterface().UpdateStatusComplete(&req)
	case constants.ONPROGRESS_DESC:
		req.Status = constants.ONPROGRESS
		res = services.InitializeOrderServiceInterface().UpdatePayment(&req)
	}

	// logs.Info(statusPayment)
	// res = services.InitializeOrderServiceInterface().Add(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res update pay order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}

// UpdateQty ...
func (controller *OrderController) UpdateQty(ctx *gin.Context) {
	fmt.Println(">>> OrderController - Update Qty <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}
	req := dto.OrderDetailRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)

	if req.ID == 0 {
		res = services.InitializeOrderServiceInterface().AddNewDetail(&req)
	} else {
		res = services.InitializeOrderServiceInterface().UpdateQty(&req)
	}

	resByte, _ := json.Marshal(res)
	log.Println("res update qty order --> ", string(resByte))

	ctx.JSON(http.StatusOK, res)

}

// GetByRestoIdTabelID ...
func (controller *OrderController) GetByRestoIdTabelID(ctx *gin.Context) {
	fmt.Println(">>> Order Controoler - Get Tabel ID <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	restoId, err := strconv.ParseInt(ctx.Param("restoId"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "Resto id not supplied")
		ctx.Abort()
		return
	}

	tabelID, err := strconv.ParseInt(ctx.Param("tabelId"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "Tabel id not supplied")
		ctx.Abort()
		return
	}

	res = services.InitializeOrderServiceInterface().GetByRestoIdTabelId(restoId, tabelID)

	ctx.JSON(http.StatusOK, res)

}

// AddItemOrderToTabel ...
func (controller *OrderController) AddItemOrderToTabel(ctx *gin.Context) {
	fmt.Println(">>> OrderController - add Item to tabel <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.AddOrderItemDto{}
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

	res = services.InitializeOrderServiceInterface().AddItemOrderToTabel(&req)
	resByte, _ := json.Marshal(res)
	log.Println("res add order --> ", string(resByte))
	ctx.JSON(http.StatusOK, res)

}

// PaymentByTabelID ...
func (controller *OrderController) PaymentByTabelID(ctx *gin.Context) {
	fmt.Println(">>> OrderController - Payment at tabel ID  <<<")
	parent := context.Background()
	defer parent.Done()

	tabelID, err := strconv.ParseInt(ctx.Param("tabelID"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "Tabel id not supplied")
		ctx.Abort()
		return
	}

	res := models.Response{}
	req := []dto.OrderPaymentDto{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	json.Marshal(req)

	res = services.InitializeOrderServiceInterface().PaymentByTabelID(req, tabelID)

	resByte, _ := json.Marshal(res)
	log.Println("res Payment", resByte)

	ctx.JSON(http.StatusOK, res)

}

// GetPaymentByTabelID ...
func (controller *OrderController) GetPaymentByTabelID(ctx *gin.Context) {
	fmt.Println(">>> OrderController - Get Payment tabel ID  <<<")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{}

	tabelID, err := strconv.ParseInt(ctx.Param("tabelID"), 10, 64)
	if err != nil {
		logs.Info("error", err)
		ctx.JSON(http.StatusBadRequest, "Tabel id not supplied")
		ctx.Abort()
		return
	}
	res = services.InitializeOrderServiceInterface().GetPaymentByTabelID(tabelID)

	resByte, _ := json.Marshal(res)
	log.Println("res Payment", resByte)

	ctx.JSON(http.StatusOK, res)

}
