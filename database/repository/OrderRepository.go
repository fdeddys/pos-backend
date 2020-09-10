package repository

import (
	"fmt"
	"log"
	"resto-be/constants"
	"resto-be/database"
	"resto-be/models"
	"resto-be/models/dbmodels"
	"resto-be/models/dto"
	// "github.com/astaxie/beego/logs"
)

func AddOrder(order *dbmodels.Order) error {
	db := database.GetDbCon()
	err := db.Save(&order).Error

	return err
}

func AddOrderDetail(orderDetail *dbmodels.OrderDetail) error {
	db := database.GetDbCon()
	err := db.Save(&orderDetail).Error

	return err
}

func GetByCustomerIdPage(req dto.OrderRequestDto, page int, limit int) ([]dbmodels.Order, error) {
	db := database.GetDbCon()

	var orders []dbmodels.Order

	log.Println("limit", limit)
	log.Println("offset", (page-1)*limit)

	if err := db.Preload("Voucher").Preload("Customer").Preload("User").Preload("Resto").Order("id desc").Limit(limit).Offset((page-1)*limit).Where("customer_Id = ?", req.CustomerId).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

// GetOrderById ...
func GetOrderById(id int64) (dbmodels.Order, error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.Order{}

	err := db.Preload("Voucher").Preload("Customer").Preload("Resto").Where(" id = ?  ", id).First(&order).Error

	return order, err

}

// GetOrderDetailByOrderID ...
func GetOrderDetailByOrderID(orderID int64) []dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []dbmodels.OrderDetail

	db.Preload("MenuItem").Preload("MenuItem.Category").Preload("MenuItem.MenuGroup").Find(&orderDetails, " order_Id = ? and qty > 0 ", orderID)

	return orderDetails
}

func GetOrderDetailByID(orderDetailID int64) dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail

	db.Preload("MenuItem").Preload("MenuItem.Category").Find(&orderDetail, " id = ? ", orderDetailID)

	return orderDetail
}

// GetByRestoCodePage ...
func GetByRestoCodePage(req dto.OrderRequestDto, page int, limit int) ([]dbmodels.Order, error) {
	db := database.GetDbCon()

	var orders []dbmodels.Order

	resto, err := GetRestoBycode(req.RestoCode)
	if err != nil {
		return orders, err
	}

	query := fmt.Sprintf("where resto_id = %d ", resto.ID)

	if req.StartDate != "" {
		startDate := req.StartDate + " 00:00:00"
		query = fmt.Sprintf("%v AND order_date >= '%v' ", query, startDate)
	}

	if req.EndDate != "" {
		endDate := req.EndDate + " 23:59:59"
		query = fmt.Sprintf("%v AND order_date <= '%v' ", query, endDate)
	}

	switch req.PaymentStatus {
	case constants.PAID_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.PAID)
	case constants.UNPAID_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.UNPAID)
	case constants.CANCEL_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.CANCEL)
	}

	log.Println("query --> ", query)

	if err := db.Preload("Customer").Preload("User").Preload("Resto").Order("id desc").Limit(limit).Offset((page - 1) * limit).Raw(" select * from public.order " + query).Find(&orders).Error; err != nil {
		return orders, err
	}
	//
	//
	//if err := db.Preload("Resto").Order("id desc").Limit(limit).Offset((page-1)*limit).Where("resto_Id = ? and order_date BETWEEN 'cas' AND 'cascsa'", req.RestoId).Find(&orders).Error; err != nil {
	//	return orders, err
	//}

	return orders, nil
}

// GetByRestoIDPage ...
func GetByRestoIDPage(req dto.OrderRequestDto, page int, limit int) ([]dbmodels.Order, int, error) {
	db := database.GetDbCon()

	var orders []dbmodels.Order
	var total int

	query := fmt.Sprintf("where resto_id = %d ", req.RestoId)

	if req.StartDate != "" {
		startDate := req.StartDate + " 00:00:00"
		query = fmt.Sprintf("%v AND order_date >= '%v' ", query, startDate)
	}

	if req.EndDate != "" {
		endDate := req.EndDate + " 23:59:59"
		query = fmt.Sprintf("%v AND order_date <= '%v' ", query, endDate)
	}

	switch req.PaymentStatus {
	case constants.PAID_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.PAID)
	case constants.UNPAID_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.UNPAID)
	case constants.CANCEL_DESC:
		query = fmt.Sprintf("%v AND is_paid = '%v'", query, constants.CANCEL)
	}

	log.Println("query --> ", query)

	if err := db.Preload("Voucher").Preload("OrderDetail").Preload("OrderDetail.MenuItem").Preload("OrderDetail.MenuItem.Category").Preload("Customer").Preload("User").Preload("Resto").Order("id desc").Limit(limit).Offset((page - 1) * limit).Raw(" select * from public.order " + query).Find(&orders).Error; err != nil {
		return orders, 0, err
	}
	if err := db.Raw("select count(*) from public.order " + query).Count(&total).Error; err != nil {
		return orders, 0, err
	}

	return orders, total, nil
}

// UpdateStatusCompleteOrder ...
func UpdateStatusCompleteOrder(orderID int64, status string) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var order dbmodels.Order

	err := db.Model(&order).Where(" Id = ?", orderID).Update("IsComplete", status).Error

	return err
}

// UpdatePayment ...
func UpdatePayment(orderID int64, statusPay string) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var order dbmodels.Order

	err := db.Model(&order).Where(" Id = ?", orderID).Update("IsPaid", statusPay).Error

	return err
}

func UpdateCookStatus(orderDetailID int64, status string) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail

	err := db.Model(&orderDetail).Where(" Id = ?", orderDetailID).Update("status", status).Error

	return err
}

// UpdateQty ...
func UpdateQty(orderDetailID int64, qty int) (dbmodels.OrderDetail, error) {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail
	// logs.Info("exec ", orderDetailID, "qty ", qty)
	err := db.Model(&orderDetail).Where("id = ?", orderDetailID).Update("qty", qty).Error

	if err != nil {
		return orderDetail, err
	}

	return orderDetail, err
}

// UpdateTotal ...
func UpdateTotal(orderID int64, subTotal int64, disc float64, total int64, tax float64, serviceCharge float64, grandTotal float64) error {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var order dbmodels.Order

	err := db.Model(&order).Where("id = ?", orderID).Update(dbmodels.Order{
		SubTotal:      subTotal,
		Disc:          int64(disc),
		Total:         total,
		Tax:           int64(tax),
		ServiceCharge: int64(serviceCharge),
		GrandTotal:    int64(grandTotal),
	}).Error

	return err
}

// GetOrderDetailByOrderDetailID ...
func GetOrderDetailByOrderDetailID(orderDetailID int64) dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetail dbmodels.OrderDetail

	db.Preload("MenuItem").Find(&orderDetail, " Id = ?", orderDetailID)

	return orderDetail
}

// GetOrderDetailReport ...
func GetOrderDetailReport(req dto.OrderRequestDto) ([]models.OrderDetailReport, error) {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []models.OrderDetailReport

	err := db.Table("order_detail a").Select("a.*,b.order_no, b.order_date, b.status order_status, c.name customer, b.is_paid, b.is_complete, b.notes, d.name menu_item,  b.grand_total").
		Joins("left join public.order b on a.order_id= b.id").
		Joins("left join public.customer c on b.customer_id = c.id").
		Joins("left join public.e_menu_item d on d.id = a.e_menu_item").
		Where("b.order_date BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Where("b.resto_id = ?", req.RestoId).
		Order("b.id desc").
		Scan(&orderDetails).Error

	return orderDetails, err
}

// GetOrderByRestoIdTabelID ...
func GetOrderByRestoIdTabelID(restoId, tabelId int64) (dbmodels.Order, error) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)
	order := dbmodels.Order{}

	err := db.Preload("Voucher").Preload("Customer").Preload("Resto").Preload("OrderDetail").Where(" table_id = ?  ", tabelId).Where(" resto_id = ? ", restoId).First(&order).Error

	return order, err

}

func GetOrderDetailByOrderIDAndItemID(orderID, itemID int64) []dbmodels.OrderDetail {

	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var orderDetails []dbmodels.OrderDetail

	db.Preload("MenuItem").Preload("MenuItem.Category").Preload("MenuItem.MenuGroup").Find(&orderDetails, " order_Id = ? and qty > 0 and e_menu_item = ? ", orderID, itemID)

	return orderDetails
}
