package services

import (
	"fmt"
	"github.com/rs/xid"
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
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

	log.Println("reqq ->", req)
	users, err := repository.GetByCustomerIdPage(*req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)
	log.Println("result : ", users)

	for i := 0; i < len(users); i++ {
		users[i].IsPaidDesc = service.GetStatusOrder(users[i].IsPaid)
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = users

	return res

}

func (service *OrderServiceInterface) GetStatusOrder(status string) string {
	switch status {
	case constants.PAID:
		return constants.PAID_DESC
	case constants.NOT_YET_PAID:
		return constants.NOT_YET_PAID_DESC
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
		IsPaid:     constants.NOT_YET_PAID,
		OrderDate:  time.Now(),
		Notes:      reqDto.Notes,
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

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = order

	return res

}

// GetOrderDetailByOrderID ...
func (service *OrderServiceInterface) GetOrderDetailByOrderID(id int64) models.Response {
	var res models.Response

	order := repository.GetOrderDetailByOrderID(id)
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

// GetByRestoPage ...
func (service *OrderServiceInterface) GetByFilterPaging(req *dto.OrderRequestDto, page int, count int) models.Response {
	var res models.Response

	log.Println("reqq ->", req)
	users, err := repository.GetByRestoIDPage(*req, page, count)
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)
	log.Println("result : ", users)

	for i := 0; i < len(users); i++ {
		users[i].IsPaidDesc = service.GetStatusOrder(users[i].IsPaid)
	}

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = users

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
	orders := repository.GetOrderDetailByOrderID(orderID)

	var total int64
	total = 0
	if len(orders) > 0 {
		for _, order := range orders {
			total = total + (int64(order.Qty) * int64(order.Price))
		}
	}

	repository.UpdateTotal(orderID, total)
}
