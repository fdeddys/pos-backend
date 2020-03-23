package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"resto-be/constants"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"
	"resto-be/models"
	"resto-be/models/dto"
	"github.com/360EntSecGroup-Skylar/excelize"
	"resto-be/utils"
	"strconv"
)

type ReportServiceInterface struct {
}

func InitializeReportServiceInterface() *ReportServiceInterface {
	return &ReportServiceInterface{}
}

var (
	pathReport string
)

func init() {
	pathReport = utils.GetEnv("PATH_REPORT", "/opt/report/")
}

func (service *ReportServiceInterface) Order(req *dto.OrderRequestDto) (models.Response, string) {

	var res models.Response

	restoId := dto.CurrRestoID
	log.Println("restoId -> ", restoId)

	req.RestoId = restoId
	orders,_, err := repository.GetByRestoIDPage(*req, 1, 99999999)

	ordersByte,_ := json.Marshal(orders)
	log.Println(string(ordersByte))

	//res.Data = orders
	//return res, ""
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res, ""
	}

	fileName := service.GenerateXlsxOrder(restoId, orders, *req)
	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res, fileName
}

func (service *ReportServiceInterface) OrderDetail(req *dto.OrderRequestDto) (models.Response, string) {

	var res models.Response

	restoId := dto.CurrRestoID
	log.Println("restoId -> ", restoId)

	req.RestoId = restoId
	orderDetails, err := repository.GetOrderDetailReport(*req)

	orderDetailsByte,_ := json.Marshal(orderDetails)
	log.Println(string(orderDetailsByte))

	//res.Data = orders
	//return res, ""
	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res, ""
	}

	fileName := service.GenerateXlsxOrderDetail(orderDetails, *req)
	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG

	return res, fileName
}

func (service *ReportServiceInterface) GenerateXlsxOrderDetail (orderDetails []models.OrderDetailReport, req dto.OrderRequestDto) string {
	log.Println("Generate Xlsx")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.

	row:= 0

	f.SetCellValue("Sheet1", "B4", "No")
	f.SetCellValue("Sheet1", "C4", "Order No")
	f.SetCellValue("Sheet1", "D4", "Tanggal Transaksi")
	f.SetCellValue("Sheet1", "E4", "Customer")
	f.SetCellValue("Sheet1", "F4", "Status Pembayaran")
	f.SetCellValue("Sheet1", "G4", "Status Pemesanan")
	f.SetCellValue("Sheet1", "H4", "Notes")
	f.SetCellValue("Sheet1", "I4", "ID Menu")
	f.SetCellValue("Sheet1", "J4", "Nama Menu")
	f.SetCellValue("Sheet1", "K4", "Qty")
	f.SetCellValue("Sheet1", "L4", "Harga")
	f.SetCellValue("Sheet1", "M4", "Grand Total")
	for i, order := range orderDetails {

		switch order.IsPaid {
		case constants.PAID:
			order.IsPaidDesc = constants.PAID_DESC
		case constants.CANCEL:
			order.IsPaidDesc = constants.CANCEL_DESC
		case constants.UNPAID:
			order.IsPaidDesc = constants.UNPAID_DESC
		}

		switch order.OrderStatus {
		case constants.ORDER_STATUS_DIPESAN:
			order.StatusDesc = constants.ORDER_STATUS_DIPESAN_DESC
		case constants.ORDER_STATUS_DIMASAK:
			order.StatusDesc = constants.ORDER_STATUS_DIMASAK_DESC
		case constants.ORDER_STATUS_DIANTAR:
			order.StatusDesc = constants.ORDER_STATUS_DIANTAR_DESC
		case constants.ORDER_STATUS_DIMEJA:
			order.StatusDesc = constants.ORDER_STATUS_DIMEJA_DESC

		}


		row = i + 5

		no := fmt.Sprintf("B%v", row)
		orderNo := fmt.Sprintf("C%v", row)
		orderDate := fmt.Sprintf("D%v", row)
		customer := fmt.Sprintf("E%v", row)
		isPaid := fmt.Sprintf("F%v", row)
		status := fmt.Sprintf("G%v", row)
		notes := fmt.Sprintf("H%v", row)
		idMenu := fmt.Sprintf("I%v", row)
		namaMenu := fmt.Sprintf("J%v", row)
		qty := fmt.Sprintf("K%v", row)
		harga := fmt.Sprintf("L%v", row)
		grandTotal := fmt.Sprintf("M%v", row)

		f.SetCellValue("Sheet1", orderNo, order.OrderNo)
		f.SetCellValue("Sheet1", no, i+1)
		f.SetCellValue("Sheet1", orderDate, utils.ConvertTime(order.OrderDate))
		f.SetCellValue("Sheet1", customer, order.Customer)
		f.SetCellValue("Sheet1", isPaid, order.IsPaidDesc)
		f.SetCellValue("Sheet1", status, order.StatusDesc)
		f.SetCellValue("Sheet1", notes, order.Notes)
		f.SetCellValue("Sheet1", idMenu, order.EMenuItem)
		f.SetCellValue("Sheet1", namaMenu, order.MenuItem)
		f.SetCellValue("Sheet1", qty, order.Qty)
		f.SetCellValue("Sheet1", harga, order.Price)
		f.SetCellValue("Sheet1", grandTotal, order.GrandTotal)

	}

	//f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	reportName := fmt.Sprintf("%v-order-detail", strconv.Itoa(int(req.RestoId)))

	if _, err := os.Stat(pathReport); err != nil {
		fmt.Println("create new folder")
		errMkdir:= os.MkdirAll(pathReport, os.ModePerm)
		log.Println(errMkdir)
	}
	fileName := fmt.Sprintf("%v%v.xlsx", pathReport,reportName)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

	return fileName
}

func (service *ReportServiceInterface) GenerateXlsxOrder (restoId int64, orders []dbmodels.Order, req dto.OrderRequestDto) string {
	log.Println("Generate Xlsx")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.

	totalDibayar := int64(0)
	totalCancel := int64(0)
	row:= 0

	f.MergeCell("Sheet1", "A1", "G1")
	f.SetCellValue("Sheet1", "A1", "LAPORAN PENJUALAN")
	f.SetCellValue("Sheet1", "A2", "From")
	f.SetCellValue("Sheet1", "B2", req.StartDate)
	f.SetCellValue("Sheet1", "A3", "To")
	f.SetCellValue("Sheet1", "B3", req.EndDate)

	f.SetCellValue("Sheet1", "A4", "No")
	f.SetCellValue("Sheet1", "B4", "Order No")
	f.SetCellValue("Sheet1", "C4", "Tanggal Transaksi")
	f.SetCellValue("Sheet1", "D4", "Customer")
	f.SetCellValue("Sheet1", "E4", "Status Pembayaran")
	f.SetCellValue("Sheet1", "F4", "Status Pemesanan")
	f.SetCellValue("Sheet1", "G4", "Notes")
	f.SetCellValue("Sheet1", "H4", "Total (Rp)")
	for i, order := range orders {

		switch order.IsPaid {
		case constants.PAID:
			order.IsPaidDesc = constants.PAID_DESC
			totalDibayar = order.Total + totalDibayar
		case constants.CANCEL:
			order.IsPaidDesc = constants.CANCEL_DESC
			totalCancel = order.Total + totalCancel
		case constants.UNPAID:
			order.IsPaidDesc = constants.UNPAID_DESC

		}

		switch order.Status {
		case constants.ORDER_STATUS_DIPESAN:
			order.StatusDesc = constants.ORDER_STATUS_DIPESAN_DESC
		case constants.ORDER_STATUS_DIMASAK:
			order.StatusDesc = constants.ORDER_STATUS_DIMASAK_DESC
		case constants.ORDER_STATUS_DIANTAR:
			order.StatusDesc = constants.ORDER_STATUS_DIANTAR_DESC
		case constants.ORDER_STATUS_DIMEJA:
			order.StatusDesc = constants.ORDER_STATUS_DIMEJA_DESC

		}


		row = i + 5

		no := fmt.Sprintf("A%v", row)
		orderNo := fmt.Sprintf("B%v", row)
		orderDate := fmt.Sprintf("C%v", row)
		customer := fmt.Sprintf("D%v", row)
		isPaid := fmt.Sprintf("E%v", row)
		status := fmt.Sprintf("F%v", row)
		notes := fmt.Sprintf("G%v", row)
		total := fmt.Sprintf("H%v", row)

		f.SetCellValue("Sheet1", orderNo, order.OrderNo)
		f.SetCellValue("Sheet1", no, i+1)
		f.SetCellValue("Sheet1", orderDate, utils.ConvertTime(order.OrderDate))
		f.SetCellValue("Sheet1", customer, order.Customer.Name)
		f.SetCellValue("Sheet1", notes, order.Notes)
		f.SetCellValue("Sheet1", isPaid, order.IsPaidDesc)
		f.SetCellValue("Sheet1", status, order.StatusDesc)
		f.SetCellValue("Sheet1", total, order.Total)

	}
	row = row +1
	start := fmt.Sprintf("A%v", row)
	end := fmt.Sprintf("G%v", row)
	dibayar := fmt.Sprintf("H%v", row)

	f.MergeCell("Sheet1", start, end)
	f.SetCellValue("Sheet1", start, "Total Penjualan Berhasil (Rp)")
	f.SetCellValue("Sheet1", dibayar, totalDibayar)

	start = fmt.Sprintf("A%v", row+1)
	end = fmt.Sprintf("G%v", row+1)
	cancel := fmt.Sprintf("H%v", row+1)

	f.MergeCell("Sheet1", start, end)
	f.SetCellValue("Sheet1", start, "Total Penjualan Batal (Rp)")
	f.SetCellValue("Sheet1", cancel, totalCancel)


	//f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	restoIdString := strconv.Itoa(int(restoId))

	if _, err := os.Stat(pathReport); err != nil {
		fmt.Println("create new folder")
		errMkdir:= os.MkdirAll(pathReport, os.ModePerm)
		log.Println(errMkdir)
	}
	fileName := fmt.Sprintf("%v%v.xlsx", pathReport,restoIdString)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

	return fileName
}