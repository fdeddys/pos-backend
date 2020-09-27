package services

import (
	"fmt"
	"log"
	"resto-be/constants"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
	"time"

	"github.com/rs/xid"

	"github.com/astaxie/beego/logs"
)

type OrderServiceInterface struct {
}

func InitializeOrderServiceInterface() *OrderServiceInterface {
	return &OrderServiceInterface{}
}

// GetByCustomerPage ...
func (service *OrderServiceInterface) GetByCustomerPage(req *dto.OrderRequestDto, page int, count int) models.Response {
	var res models.Response

	customer, err := repository.GetCustomerByID(req.CustomerId)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	if customer.ManualRestoID > 0 {
		page = 1
		count = 1
	}
	orders, err := repository.GetByCustomerIdPage(*req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)
	log.Println("result : ", orders)

	for i := 0; i < len(orders); i++ {
		orders[i].IsPaidDesc = service.GetStatusOrder(orders[i].IsPaid)
		orders[i].IsCompleteDesc = service.GetStatusComplete(orders[i].IsComplete)

		switch orders[i].Status {
		case constants.ORDER_STATUS_DIPESAN:
			orders[i].StatusDesc = constants.ORDER_STATUS_DIPESAN_DESC
		case constants.ORDER_STATUS_DIMASAK:
			orders[i].StatusDesc = constants.ORDER_STATUS_DIMASAK_DESC
		case constants.ORDER_STATUS_DIANTAR:
			orders[i].StatusDesc = constants.ORDER_STATUS_DIANTAR_DESC
		case constants.ORDER_STATUS_DIMEJA:
			orders[i].StatusDesc = constants.ORDER_STATUS_DIMEJA_DESC
		}

	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orders

	return res

}

func (service *OrderServiceInterface) GetStatusComplete(status string) string {
	switch status {
	case constants.COMPLETE:
		return constants.COMPLETE_DESC
	case constants.ONPROGRESS:
		return constants.ONPROGRESS_DESC
	}

	return "-"
}

func (service *OrderServiceInterface) GetStatusOrder(status string) string {
	switch status {
	case constants.PAID:
		return constants.PAID_DESC
	case constants.UNPAID:
		return constants.UNPAID_DESC
	case constants.CANCEL:
		return constants.CANCEL_DESC
	}

	return "-"
}

func (service *OrderServiceInterface) GenerateOrderNumber(restoId int64) string {

	return generateOrderNo(restoId)
}

func generateOrderNo(restoID int64) string {
	resto, _ := repository.GetRestoById(restoID)
	x := xid.New().Counter()

	orderNo := fmt.Sprintf("%v%v", resto.RestoCode, x)
	return orderNo

}

func (service *OrderServiceInterface) Add(reqDto *dto.OrderRequestDto) models.Response {
	var res models.Response

	orderNo := service.GenerateOrderNumber(reqDto.RestoId)

	/*pack message order*/
	order := dbmodels.Order{
		OrderNo:     orderNo,
		TableId:     reqDto.TableId,
		RestoId:     reqDto.RestoId,
		CustomerId:  reqDto.CustomerId,
		Total:       reqDto.Total,
		UserId:      dto.CurrUserID,
		Status:      constants.ORDER_STATUS_DIPESAN,
		IsPaid:      constants.UNPAID,
		OrderDate:   time.Now(),
		Notes:       reqDto.Notes,
		Disc:        reqDto.Disc,
		VoucherCode: reqDto.VoucherCode,
	}

	// save order to db
	err := repository.AddOrder(&order)
	if err != nil {
		log.Println("err save database : ", err)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}

	// saver order detail
	orderId := order.ID

	res = service.AddOrderDetail(orderId, reqDto.OrderDetails)
	reCalculate(orderId)

	return res
}

func (service *OrderServiceInterface) AddOrderDetail(orderId int64, orderDetails []dto.OrderDetailRequest) models.Response {
	var res models.Response

	for _, detail := range orderDetails {
		log.Println(detail)
		// get menuItem by ID
		menuItem, err := repository.GetMenuItemById(detail.EMenuItem)
		if err != nil {
			log.Println("err get menu item from database : ", err)

			res.Rc = constants.ERR_CODE_11
			res.Msg = constants.ERR_CODE_11_MSG
			return res
		}

		// pack msg order detail
		orderDetail := dbmodels.OrderDetail{
			Price:     menuItem.Price,
			EMenuItem: detail.EMenuItem,
			Qty:       detail.Qty,
			OrderId:   orderId,
			Status:    constants.COOK_WAITING,
		}

		// save order detail to db
		errOrderDetail := repository.AddOrderDetail(&orderDetail)
		if err != nil {
			log.Println("err save orderdetail to database : ", errOrderDetail)

			res.Rc = constants.ERR_CODE_10
			res.Msg = constants.ERR_CODE_10_MSG
			return res
		}

		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = fmt.Sprintf("%v", orderId)

	return res
}

// GetById ...
func (service *OrderServiceInterface) GetById(id int64) models.Response {
	var res models.Response

	order, err := repository.GetOrderById(id)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	order.IsPaidDesc = service.GetStatusOrder(order.IsPaid)
	order.IsCompleteDesc = service.GetStatusComplete(order.IsComplete)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = order

	return res

}

// GetOrderDetailByOrderID ...
func (service *OrderServiceInterface) GetOrderDetailByOrderID(id int64) models.Response {
	var res models.Response

	order := repository.GetOrderDetailByOrderID(id)

	for i := 0; i < len(order); i++ {
		switch order[i].Status {
		case constants.COOK_WAITING:
			order[i].StatusDesc = constants.COOK_WAITING_DESC
		case constants.COOK_COOKING:
			order[i].StatusDesc = constants.COOK_COOKING_DESC
		case constants.COOK_DELIVERY:
			order[i].StatusDesc = constants.COOK_DELIVERY_DESC
		case constants.COOK_AT_LOCATION:
			order[i].StatusDesc = constants.COOK_AT_LOCATION_DESC
		case constants.COOK_ON_HAND:
			order[i].StatusDesc = constants.COOK_ON_HAND_DESC
		case constants.COOK_CANCEL:
			order[i].StatusDesc = constants.COOK_CANCEL_DESC

		}
	}

	// if err != nil {
	// 	log.Println("err get from database : ", err)

	// 	res.Rc = constants.ERR_CODE_11
	// 	res.Msg = constants.ERR_CODE_11_MSG
	// 	return res
	// }

	// log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = order

	return res

}

// GetByRestoPage ...
func (service *OrderServiceInterface) GetByFilterPaging(req *dto.OrderRequestDto, page int, count int) models.Response {
	var res models.Response

	log.Println("reqq ->", req)
	orders, total, err := repository.GetByRestoIDPage(*req, page, count)

	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)
	log.Println("result : ", orders)

	for i := 0; i < len(orders); i++ {
		orders[i].IsPaidDesc = service.GetStatusOrder(orders[i].IsPaid)
		orders[i].IsCompleteDesc = service.GetStatusComplete(orders[i].IsComplete)
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orders
	res.TotalData = total

	return res

}

func (service *OrderServiceInterface) UpdateCookStatus(req *dto.OrderRequestDto) models.Response {
	var res models.Response

	orderDetailID := req.ID
	orderDetail := repository.GetOrderDetailByID(orderDetailID)

	if orderDetail.ID == 0 {
		log.Println("data not found")

		res.Rc = constants.ERR_CODE_30
		res.Msg = constants.ERR_CODE_30_MSG
		return res
	}

	if errUpdate := repository.UpdateCookStatus(orderDetailID, req.Status); errUpdate != nil {
		res.Rc = constants.ERR_CODE_13
		res.Msg = constants.ERR_CODE_13_MSG
		res.Data = nil
		return res
	}

	switch req.Status {
	case constants.COOK_WAITING:
		orderDetail.Status = constants.COOK_WAITING
		orderDetail.StatusDesc = constants.COOK_WAITING_DESC
	case constants.COOK_COOKING:
		orderDetail.Status = constants.COOK_COOKING
		orderDetail.StatusDesc = constants.COOK_COOKING_DESC
	case constants.COOK_DELIVERY:
		orderDetail.Status = constants.COOK_DELIVERY
		orderDetail.StatusDesc = constants.COOK_DELIVERY_DESC
	case constants.COOK_AT_LOCATION:
		orderDetail.Status = constants.COOK_AT_LOCATION
		orderDetail.StatusDesc = constants.COOK_AT_LOCATION_DESC
	case constants.COOK_ON_HAND:
		orderDetail.Status = constants.COOK_ON_HAND_DESC
		orderDetail.StatusDesc = constants.COOK_ON_HAND_DESC
	case constants.COOK_CANCEL:
		orderDetail.Status = constants.COOK_CANCEL
		orderDetail.StatusDesc = constants.COOK_CANCEL_DESC

	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	res.Data = orderDetail

	return res
}

// UpdatePayment ...
func (service *OrderServiceInterface) UpdateStatusComplete(req *dto.OrderRequestDto) models.Response {
	var res models.Response

	log.Println("reqq ->", req)
	orderID := req.ID
	order, err := repository.GetOrderById(orderID)

	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	if errUpdate := repository.UpdateStatusCompleteOrder(orderID, req.Status); errUpdate != nil {
		res.Rc = constants.ERR_CODE_13
		res.Msg = constants.ERR_CODE_13_MSG
		res.Data = nil
		return res
	}

	switch req.Status {
	case constants.COMPLETE:
		order.IsComplete = constants.COMPLETE
		order.IsCompleteDesc = constants.COMPLETE_DESC
	case constants.ONPROGRESS:
		order.IsComplete = constants.ONPROGRESS
		order.IsCompleteDesc = constants.ONPROGRESS_DESC
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	res.Data = order

	return res

}

// UpdatePayment ...
func (service *OrderServiceInterface) UpdatePayment(req *dto.OrderRequestDto) models.Response {
	var res models.Response

	log.Println("reqq ->", req)
	orderID := req.ID
	order, err := repository.GetOrderById(orderID)

	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	// proses REQ paid hanya bisa ORDER dg status 00
	if req.Status == "10" && order.IsPaid != "00" {

		switch order.IsPaid {
		case constants.PAID:
			res.Rc = constants.ERR_CODE_14
			res.Msg = constants.ERR_CODE_14_MSG
		case constants.CANCEL:
			res.Rc = constants.ERR_CODE_15
			res.Msg = constants.ERR_CODE_15_MSG
		}
		return res
	}
	// proses cancel (REQ) hanya bisa status ORDER 00 atau 10
	// jika status sdh 20 reject
	if req.Status == constants.CANCEL && order.IsPaid == constants.CANCEL {
		res.Rc = constants.ERR_CODE_15
		res.Msg = constants.ERR_CODE_15_MSG
		return res
	}

	if errUpdate := repository.UpdatePayment(orderID, req.Status); errUpdate != nil {
		res.Rc = constants.ERR_CODE_13
		res.Msg = constants.ERR_CODE_13_MSG
		res.Data = nil
		return res
	}

	reCalculate(orderID)

	log.Println("req.Status -->", req.Status)

	switch req.Status {
	case constants.PAID:
		order.IsPaid = constants.PAID
		order.IsPaidDesc = constants.PAID_DESC
	case constants.CANCEL:
		order.IsPaid = constants.CANCEL
		order.IsPaidDesc = constants.CANCEL_DESC
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	res.Data = order

	return res

}

// UpdateQty ...
func (service *OrderServiceInterface) UpdateQty(req *dto.OrderDetailRequest) models.Response {
	var res models.Response

	logs.Info("Update detail", req)
	detail, errUpdate := repository.UpdateQty(req.ID, req.Qty)
	if errUpdate != nil {
		res.Rc = constants.ERR_CODE_13
		res.Msg = constants.ERR_CODE_13_MSG
		res.Data = nil
		return res
	}
	logs.Info("isi detail ", detail)

	orderDetail := repository.GetOrderDetailByOrderDetailID(req.ID)
	reCalculate(orderDetail.OrderId)
	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orderDetail.ID
	res.TotalData = 0
	return res

}

func reCalculate(orderID int64) {
	order, _ := repository.GetOrderById(orderID)
	orders := repository.GetOrderDetailByOrderID(orderID)
	resto, _ := repository.GetRestoById(order.RestoId)

	var subTotal int64
	subTotal = 0
	if len(orders) > 0 {
		for _, order := range orders {
			subTotal = subTotal + (int64(order.Qty) * int64(order.Price))
		}
	}
	log.Println("voucher ", order.VoucherCode)
	log.Println("subTotal -->", subTotal)

	// validasi voucher
	reqVoucher := dto.VoucherRequestDto{
		RestoId: order.RestoId,
		Code:    order.VoucherCode,
	}

	voucher, err := repository.GetVoucherByCode(reqVoucher)
	if err != nil {
		log.Println("err coucher ", err)
	}
	log.Println("voucher --> ", voucher)
	var disc float64
	maxValue := float64(voucher.MaxValue)
	voucherValue := float64(voucher.Value)
	if subTotal >= voucher.MinPayment {
		disc = float64(subTotal) * voucherValue / 100
		log.Println("disc ", disc)
		if disc > maxValue {
			disc = maxValue
		}
	}

	total := subTotal - int64(disc)
	tax := (float64(resto.Tax) / 100) * float64(total)
	serviceCharge := (float64(resto.ServiceCharge) / 100) * float64(total)

	grandTotal := float64(total) + tax + serviceCharge
	log.Println("total --> ", total)
	log.Println("serviceCharge --> ", serviceCharge)
	log.Println("tax --> ", tax)
	log.Println("grandTotal --> ", grandTotal)

	repository.UpdateTotal(orderID, subTotal, disc, total, tax, serviceCharge, grandTotal)
}

// GetByTabelId ...
func (service *OrderServiceInterface) GetByRestoIdTabelId(restoId, tabelid int64) models.Response {
	var res models.Response

	log.Println("search order ID from tabel id [", tabelid, "] ")
	// order, err := repository.GetOrderByRestoIdTabelID(restoId, tabelid)
	orderID, err := repository.GetOrderIdFromTabelID(tabelid)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG + " [tabel not found]"
		return res
	}

	if orderID == 0 {
		log.Println("No order ID found for this tabel [ table not occupied ] ")
		res.Rc = constants.ERR_CODE_70
		res.Msg = constants.ERR_CODE_70_MSG
		return res
	}

	order, errOrder := repository.GetOrderById(orderID)
	if errOrder != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG + " [Order invalid]"
		return res
	}

	log.Println("get data : ", res)

	order.IsPaidDesc = service.GetStatusOrder(order.IsPaid)
	order.IsCompleteDesc = service.GetStatusComplete(order.IsComplete)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = order

	return res

}

// AddItemOrderToTabel ...
func (service *OrderServiceInterface) AddItemOrderToTabel(req *dto.AddOrderItemDto) models.Response {
	var res models.Response
	var orderID int64
	var msg string
	var tabel dbmodels.Table
	var errSaveHd string

	log.Println("Order service - add Item Order to Tabel")
	orderID, msg, tabel = getOrderIdByTabel(req.TableID)
	if orderID == -1 {
		res.Rc = constants.ERR_CODE_10
		res.Msg = msg
		return res
	}

	if orderID == 0 {
		orderID, errSaveHd = saveOrderAndUpdateToTabel(tabel)
		if orderID == 0 {
			res.Rc = constants.ERR_CODE_10
			res.Msg = errSaveHd
			res.Data = orderID
		}
	}

	orderDetail := packOrderDetail(orderID, req.ItemID)
	if errAddItem := repository.AddOrderDetail(&orderDetail); errAddItem != nil {
		res.Rc = constants.ERR_CODE_10
		res.Msg = errAddItem.Error()
		res.Data = orderID

		return res
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orderID

	return res

}

func getOrderIdByTabel(tabelID int64) (orderID int64, msg string, tabel dbmodels.Table) {

	fmt.Println(">>> Validate request add item to order <<<")

	orderID = -1
	msg = ""

	tabel, errTabel := repository.GetTabelByTableID(tabelID)
	if errTabel != nil {
		return orderID, errTabel.Error(), dbmodels.Table{}
	}

	if tabel.Status == constants.TABEL_STATUS_INACTIVE {
		return orderID, "Tabel inactive", dbmodels.Table{}
	}

	orderID = tabel.OrderID

	fmt.Println(">>> END Validate request add item to order <<<")
	return orderID, "OK", tabel
}

func packOrderDetail(orderID, itemID int64) dbmodels.OrderDetail {
	var orderDetail dbmodels.OrderDetail

	// 1. cek apakah item sudah ada di order detail, jika ada update exiting qty
	// 2. - jika tidak ada create order detail id = 0

	item, _ := repository.GetMenuItemById(itemID)

	orderDetails := repository.GetOrderDetailByOrderIDAndItemID(orderID, itemID)
	if len(orderDetails) > 0 {
		orderDetail = orderDetails[0]
		orderDetail.Qty++
		return orderDetail
	}

	orderDetail.ID = 0
	orderDetail.EMenuItem = itemID
	orderDetail.OrderId = orderID
	orderDetail.Price = item.Price
	orderDetail.Qty = 1
	orderDetail.Status = constants.COOK_WAITING

	return orderDetail
}

func saveOrderAndUpdateToTabel(tabel dbmodels.Table) (int64, string) {

	order := packOrder(tabel)

	if errSave := repository.AddOrder(&order); errSave != nil {
		return 0, "Failed save Order HDR >> " + errSave.Error()
	}

	if errUpdate := repository.UpdateOrderIdToTabel(tabel.ID, order.ID); errUpdate != nil {
		return 0, "Failed save order ID to Table >> " + errUpdate.Error()
	}
	return order.ID, ""
}

func packOrder(tabel dbmodels.Table) dbmodels.Order {
	var order dbmodels.Order

	fmt.Println("tabel          => ", tabel.ID)
	fmt.Println("group table    => ", tabel.GroupTable.ID)
	fmt.Println("resto by group => ", tabel.GroupTable.RestoCode)

	resto, _ := repository.GetRestoByRestoCode(tabel.GroupTable.RestoCode)
	fmt.Println("Resto          => ", resto.Name)

	orderNo := generateOrderNo(resto.ID)

	order.Disc = 0
	order.GrandTotal = 0
	order.IsComplete = "00"
	order.IsPaid = constants.UNPAID
	order.OrderDate = time.Now()
	order.OrderNo = orderNo
	order.RestoId = resto.ID
	order.ServiceCharge = int64(resto.ServiceCharge)
	order.Status = constants.ORDER_STATUS_DIPESAN
	order.SubTotal = 0
	order.TableId = tabel.ID
	order.Tax = int64(resto.Tax)
	order.Total = 0
	order.UserId = dto.CurrUserID

	return order
}

// PaymentByTabelID ...
func (service *OrderServiceInterface) PaymentByTabelID(req []dto.OrderPaymentDto, tabelID int64) models.Response {
	var res models.Response

	log.Println("Order service - Find Order ID ")

	// check tabel isValid
	orderID, msg, tabel := getOrderIdByTabel(tabelID)
	if orderID == -1 {
		res.Rc = constants.ERR_CODE_10
		res.Msg = msg
		return res
	}

	if tabel.Status != 20 {
		res.Rc = constants.ERR_CODE_10
		res.Msg = "Tabel tidak sedang ada order konsumen / Tabel Empty !"
		return res
	}

	orderID = tabel.OrderID
	// if !validatePaymentTotal(orderID, req) {
	// 	res.Rc = constants.ERR_CODE_10
	// 	res.Msg = "Total Order lebih besar dari payment, please refresh data !"
	// 	return res
	// }

	// update Jumlah Pembayaran per Type mis cash Rp xxx, atau debit ro xxx
	for _, orderPayment := range req {
		if err := repository.UpdateAmountById(orderPayment.ID, orderPayment.Total); err != nil {
			fmt.Println("error ", err.Error())
		}
		// orderPayment.ID
	}
	repository.UpdatePayment(orderID, "10")
	repository.ReleaseTabel(tabelID)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orderID

	return res

}

func validatePaymentTotal(orderID int64, orderPayments []dto.OrderPaymentDto) bool {

	orderDetails := repository.GetOrderDetailByOrderID(orderID)
	totalOrder := float64(0)
	for _, orderDetail := range orderDetails {
		totalOrder += float64(orderDetail.Qty) * orderDetail.Price
	}

	totalPayment := float64(0)
	for _, orderPayment := range orderPayments {
		totalPayment += orderPayment.Total
		// orderPayment.ID
	}

	return (totalPayment >= totalOrder)

}

// GetPaymentByTabelID ...
func (service *OrderServiceInterface) GetPaymentByTabelID(tabelID int64) models.Response {
	var res models.Response
	log.Println("Payment Order service - Find Order-Payment by Tabel ID ")

	// check tabel isValid
	orderID, msg, tabel := getOrderIdByTabel(tabelID)
	if orderID == -1 {
		res.Rc = constants.ERR_CODE_10
		res.Msg = msg
		return res
	}

	orderID = tabel.OrderID
	orderPayments := repository.GetOrderPaymentByOrderId(orderID)

	if len(orderPayments) < 1 {

		order, _ := repository.GetOrderById(orderID)
		paymentTypes, _ := repository.GetPaymentTypeByRestoID(order.RestoId)
		if len(paymentTypes) < 1 {
			res.Rc = constants.ERR_CODE_11
			res.Msg = "Payment type for this resto not yet setting, check [ PAYMENT_TYPES ] !!"
			res.Data = orderPayments
		}
		for _, paymentType := range paymentTypes {
			var orderPayment dbmodels.OrderPayment
			orderPayment.OrderID = orderID
			orderPayment.PaymentTypeId = paymentType.ID
			orderPayment.PaymentTypeName = paymentType.Name
			orderPayment.Total = 0
			repository.SaveOrderPayment(&orderPayment)
			orderPayments = append(orderPayments, orderPayment)
		}
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orderPayments

	return res

}

// AddNewDetail ...
func (service *OrderServiceInterface) AddNewDetail(req *dto.OrderDetailRequest) models.Response {

	var res models.Response
	menuItem, err := repository.GetMenuItemById(req.EMenuItem)

	if err != nil {
		log.Println("err get menu item from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	// pack msg order detail
	orderDetail := dbmodels.OrderDetail{
		Price:     menuItem.Price,
		EMenuItem: req.EMenuItem,
		Qty:       req.Qty,
		OrderId:   req.OrderID,
		Status:    constants.COOK_WAITING,
	}

	// save order detail to db
	errOrderDetail := repository.AddOrderDetail(&orderDetail)
	if err != nil {
		log.Println("err save orderdetail to database : ", errOrderDetail)

		res.Rc = constants.ERR_CODE_10
		res.Msg = constants.ERR_CODE_10_MSG
		return res
	}
	reCalculate(orderDetail.OrderId)
	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = orderDetail.ID
	res.TotalData = 0
	return res

}
