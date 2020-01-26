package services

import (
	"log"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"time"
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

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = users

	return res

}

func (service *OrderServiceInterface) Add(reqDto *dto.OrderRequestDto) models.Response {
	var res models.Response

	/*pack message order*/
	order := dbmodels.Order{
		OrderNo:    reqDto.OrderNo,
		TableId:    reqDto.TableId,
		RestoId:    reqDto.RestoId,
		CustomerId: reqDto.CustomerId,
		Total:      reqDto.Total,
		UserId:     dto.CurrUserID,
		Status:     constants.ORDER_STATUS_DIPESAN,
		IsPaid:     constants.NOT_YET_PAID,
		OrderDate:  time.Now(),
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
