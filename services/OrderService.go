package services

import (
	"fmt"
	"github.com/rs/xid"
	"log"
	"resto-be/constants"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"time"

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

	resto, _ := repository.GetRestoById(restoId)
	x:= xid.New().Counter()

	orderNo := fmt.Sprintf("%v%v", resto.RestoCode, x)

	return orderNo
}

func (service *OrderServiceInterface) Add(reqDto *dto.OrderRequestDto) models.Response {
	var res models.Response

	orderNo := service.GenerateOrderNumber(reqDto.RestoId)

	/*pack message order*/
	order := dbmodels.Order{
		OrderNo:    orderNo,
		TableId:    reqDto.TableId,
		RestoId:    reqDto.RestoId,
		CustomerId: reqDto.CustomerId,
		Total:      reqDto.Total,
		UserId:     dto.CurrUserID,
		Status:     constants.ORDER_STATUS_DIPESAN,
		IsPaid:     constants.UNPAID,
		OrderDate:  time.Now(),
		Notes:      reqDto.Notes,
		Disc:		reqDto.Disc,
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
			Status: constants.COOK_WAITING,
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
	res.Msg = constants.ERR_CODE_00_MSG

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

	for i:=0; i< len(order); i++ {
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
	orders, total,  err := repository.GetByRestoIDPage(*req, page, count)

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
	res.Data = orderDetail

	return res

}

func reCalculate(orderID int64) {
	order,_ := repository.GetOrderById(orderID)
	orders := repository.GetOrderDetailByOrderID(orderID)
	resto,_ := repository.GetRestoById(order.RestoId)

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
	reqVoucher:= dto.VoucherRequestDto{
		RestoId: order.RestoId,
		Code: order.VoucherCode,
	}

	voucher, err:=repository.GetVoucherByCode(reqVoucher)
	if err!= nil {
		log.Println("err coucher ", err)
	}
	log.Println("voucher --> ", voucher)
	var disc float64
	maxValue := float64(voucher.MaxValue)
	voucherValue := float64(voucher.Value)
	if subTotal >= voucher.MinPayment {
		disc = float64(subTotal) * voucherValue/100
		log.Println("disc ", disc)
		if disc > maxValue {
			disc = maxValue
		}
	}

	total := subTotal - int64(disc)
	tax := (float64(resto.Tax)/100) * float64(total)
	serviceCharge := (float64(resto.ServiceCharge)/100) * float64(total)

	grandTotal := float64(total) + tax + serviceCharge
	log.Println("total --> ", total)
	log.Println("serviceCharge --> ", serviceCharge)
	log.Println("tax --> ", tax)
	log.Println("grandTotal --> ", grandTotal)


	repository.UpdateTotal(orderID, subTotal, disc, total, tax, serviceCharge, grandTotal)
}
