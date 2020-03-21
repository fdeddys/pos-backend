package report

import (
	// "distribution-system-be/database"
	"fmt"
	"log"
	"resto-be/models/dbmodels"
	"resto-be/database/repository"

	// database "resto-be/database/repository"

	"github.com/leekchan/accounting"
	"github.com/signintech/gopdf"
)

type DataRptHdr struct {
	RestoAddress string
	RestoName    string
	TransAt      string
	OrderNo      string
}

type DataReptDetail struct {
	Item     string
	Quantity int64
	Unit     string
	Price    int64
	Total    int64
}

var (
	dataRptHdr DataRptHdr
)

func GenerateReceiveReport(orderId int64) {

	spaceLen = 15
	pageMargin = 12

	curPage = 1

	tblCol1 = 25
	tblCol2 = 80
	tblCol3 = 300
	tblCol4 = 370
	tblCol5 = 430
	tblCol6 = 500

	spaceCustomerInfo = 300
	spaceTitik = spaceCustomerInfo + 150
	spaceValue = spaceCustomerInfo + 160

	spaceSummaryInfo = spaceCustomerInfo
	spaceTitikSumamry = spaceTitik
	spaceValueSummary = spaceValue

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.SetMargins(pageMargin, pageMargin, pageMargin, pageMargin)
	pdf.AddPage()

	if err := pdf.AddTTFFont("open-sans", "font/OpenSans-Regular.ttf"); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}

	if err := pdf.AddTTFFont("open-sans-bold", "font/OpenSans-Bold.ttf"); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}

	// untuk nomor urut barang
	number = 1

	// get Data mockup utk display ke grid
	fmt.Println("data fillData Details : ", orderId)
	dataDetails := fillDataDetail(orderId)

	fmt.Println("hasil fill")
	for i, orderDetail := range dataDetails {
		fmt.Println(i, "====", orderDetail)
	}
	fmt.Println("=============")
	// setFont(&pdf, 12)
	setHeaderOrder(&pdf)
	pdf.Br(20)

	setOrderDetail(&pdf, dataDetails)
	setSummaryReceive(&pdf)
	// setSignReceive(&pdf, "Admin", "Warehouse", "Supplier")
	// 595, H: 842
	// pdf.SetFont("open-sans", "", 14)

	// pdf.SetFont("open-sans", "", 10)
	// for i := 2; i <= 83; i++ {
	// 	pdf.SetX(1)
	// 	pdf.SetY(10 * float64(i))
	// 	pdf.Text(fmt.Sprintf("%v", i))
	// }
	pdf.WritePdf("invoice.pdf")

}

func fillDataDetail(orderId int64) []DataReptDetail {

	order, err := repository.GetOrderById(orderId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(order)

	// orderNo = order.OrderNo
	orderDetails := repository.GetOrderDetailByOrderID(orderId)

	fmt.Println("order Details : ", orderDetails)

	fillDataResto(order)
	// tdk blh kosong
	// per halaman max 25 item detail
	totalRec = len(orderDetails)
	res := make([]DataReptDetail, totalRec+1)
	var data DataReptDetail

	subTotal = 0
	tax = 0
	grandTotal = 0
	for i, orderDetail := range orderDetails {
		data.Item = orderDetail.MenuItem.Name
		data.Quantity = int64(orderDetail.Qty)
		data.Unit = "pc"
		data.Price = int64(orderDetail.Price)
		total := data.Price * data.Quantity
		data.Total = int64(orderDetail.Price) * int64(orderDetail.Qty)
		subTotal += total
		res[i+1] = data
		fmt.Println("total sub total", subTotal)
	}
	totalRec = len(res)
	fmt.Println("Jumlah record [fill] =>", totalRec)

	tax = subTotal / 10
	grandTotal = subTotal + tax

	return res
}

func fillDataResto(order dbmodels.Order) {
	dataRptHdr.RestoName = order.Resto.Name
	dataRptHdr.RestoAddress = order.Resto.Address
	dataRptHdr.TransAt = order.OrderDate.Format("02-01-2006")
	dataRptHdr.OrderNo = order.OrderNo
}

func setHeaderOrder(pdf *gopdf.GoPdf) {
	// showLogoRecv(pdf)

	showResto(pdf)
	showInvoiceLabel(pdf)
	// space(pdf)

	// showLine(pdf)
	// showOrderNo(pdf)
}

func showLogoRecv(pdf *gopdf.GoPdf) {

	imgSize := spaceLen * 5
	posX := 20.0
	posY := spaceLen

	pdf.Image("imgs/logo.jpg", posX, posY, &gopdf.Rect{W: imgSize + 68, H: imgSize})
}

// func showOrderNo(pdf *gopdf.GoPdf) {

// 	pdf.SetY(30)
// 	pdf.SetX(450)
// 	setFontBold(pdf, 10)
// 	pdf.Text("ORDER")

// 	space(pdf)
// 	setFont(pdf, 12)
// 	pdf.SetX(450)
// 	pdf.Text(orderNo)
// }

func showInvoiceLabel(pdf *gopdf.GoPdf) {

	pdf.SetY(spaceLen * 3)
	pdf.SetX(tblCol1)
	setFontBold(pdf, 14)
	pdf.Text("INVOICE")
}

func setOrderDetail(pdf *gopdf.GoPdf, data []DataReptDetail) {

	setPageNumb(pdf, curPage)
	pdf.SetX(20)
	pdf.SetY(spaceLen * 6)

	// showResto(pdf)

	space(pdf)
	showHeaderTableOrder(pdf)

	fmt.Println("Panjang array ", len(data), "] ")
	fmt.Println("Total rec => set detail => ", totalRec, "] ")
	fmt.Println("start iterate")
	// var dataDetail DataDetail
	if totalRec > 1 {
		for i := 1; i <= 25; i++ {
			fmt.Println("idx ke [", i, "]", data[number])
			space(pdf)
			showDataDetailOrder(pdf, fmt.Sprintf("%v", number), data[number].Item, data[number].Unit, data[number].Quantity, data[number].Price, data[number].Total)
			number++
			if number >= totalRec {
				break
			}
		}
	}
	// }

	space(pdf)
	showLine(pdf)

	// jika data masih ada utk next page
	// 1. add page
	// 2. set header
	// 3. rekursif
	if totalRec > number {
		fmt.Println("NEW page")
		curPage++
		pdf.AddPage()
		showHeaderTableOrder(pdf)
		setOrderDetail(pdf, data)
	}
}

func setSummaryReceive(pdf *gopdf.GoPdf) {

	rectangle := gopdf.Rect{}
	rectangle.UnitsToPoints(gopdf.Unit_PT)

	ac := accounting.Accounting{Symbol: "", Precision: 0, Thousand: ".", Decimal: ","}
	setFont(pdf, 10)

	space(pdf)
	// pdf.SetY(spaceLen * 42)

	pdf.SetX(spaceSummaryInfo)
	// pdf.Text("Subtotal")
	pdf.CellWithOption(&rectangle, "Subtotal ", gopdf.CellOption{Align: gopdf.Left, Border: 0, Float: gopdf.Left})
	pdf.SetX(spaceTitikSumamry)
	// pdf.Text(":")
	pdf.CellWithOption(&rectangle, ": ", gopdf.CellOption{Align: gopdf.Center, Border: 0, Float: gopdf.Center})
	// pdf.SetX(spaceValueSummary)
	// pdf.Text(fmt.Sprintf("%v", subTotal))
	// pdf.Text(ac.FormatMoney(subTotal))
	fmt.Println("isi space summ ", spaceValueSummary)
	pdf.SetX(spaceValueSummary + 100)
	pdf.CellWithOption(&rectangle, ac.FormatMoney(subTotal), gopdf.CellOption{Align: gopdf.Right, Border: 0, Float: gopdf.Top})

	space(pdf)
	pdf.SetX(spaceSummaryInfo)
	// pdf.Text("Tax ")
	pdf.CellWithOption(&rectangle, "Tax", gopdf.CellOption{Align: gopdf.Left, Border: 0, Float: gopdf.Left})
	pdf.SetX(spaceTitikSumamry)
	// pdf.Text(":")
	pdf.CellWithOption(&rectangle, ": ", gopdf.CellOption{Align: gopdf.Center, Border: 0, Float: gopdf.Center})
	// pdf.SetX(spaceValueSummary)
	// pdf.Text(fmt.Sprintf("%v", tax))
	// pdf.Text(ac.FormatMoney(tax))
	pdf.SetX(spaceValueSummary + 100)
	pdf.CellWithOption(&rectangle, ac.FormatMoney(tax), gopdf.CellOption{Align: gopdf.Right, Border: 0, Float: gopdf.Top})

	space(pdf)
	pdf.SetX(spaceSummaryInfo)
	// pdf.Text("GrandTotal ")
	pdf.CellWithOption(&rectangle, "GrandTotal", gopdf.CellOption{Align: gopdf.Left, Border: 0, Float: gopdf.Left})

	pdf.SetX(spaceTitikSumamry)
	// pdf.Text(":")
	pdf.CellWithOption(&rectangle, ": ", gopdf.CellOption{Align: gopdf.Center, Border: 0, Float: gopdf.Center})
	// pdf.SetX(spaceValueSummary)
	// // pdf.Text(fmt.Sprintf("%v", grandTotal))
	// pdf.Text(ac.FormatMoney(grandTotal))
	pdf.SetX(spaceValueSummary + 100)
	pdf.CellWithOption(&rectangle, ac.FormatMoney(grandTotal), gopdf.CellOption{Align: gopdf.Right, Border: 0, Float: gopdf.Top})

}

func showHeaderTableOrder(pdf *gopdf.GoPdf) {

	showLine(pdf)
	space(pdf)
	setFontBold(pdf, 10)
	pdf.SetX(tblCol1)
	pdf.Text("#")

	pdf.SetX(tblCol2)
	pdf.Text("Item")

	pdf.SetX(tblCol3)
	pdf.Text("Quantity")

	pdf.SetX(tblCol4)
	pdf.Text("Unit")

	pdf.SetX(tblCol5)
	pdf.Text("Price")

	pdf.SetX(tblCol6)
	pdf.Text("Total")

	space(pdf)
	showLine(pdf)
}

func showDataDetailOrder(pdf *gopdf.GoPdf, no, item, unit string, qty, price, total int64) {

	ac := accounting.Accounting{Symbol: "", Precision: 0, Thousand: ".", Decimal: ","}
	setFont(pdf, 10)
	pdf.SetX(tblCol1)
	pdf.Text(no)

	pdf.SetX(tblCol2)
	pdf.Text(item)

	pdf.SetX(tblCol3)
	pdf.Text(fmt.Sprintf("%v", qty))

	pdf.SetX(tblCol4)
	pdf.Text(unit)

	pdf.SetX(tblCol5)
	// pdf.Text(fmt.Sprintf("%v", price))
	pdf.Text(ac.FormatMoney(price))

	pdf.SetX(tblCol6)
	// pdf.Text(fmt.Sprintf("%v", total))
	pdf.Text(ac.FormatMoney(total))
}

func showResto(pdf *gopdf.GoPdf) {
	// , code, name, transDate, ssNo string
	space(pdf)
	setFont(pdf, 10)

	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Resto ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(dataRptHdr.RestoName)

	space(pdf)
	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Address ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(dataRptHdr.RestoAddress)

	space(pdf)
	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Transaction at ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(dataRptHdr.TransAt)

	space(pdf)

	pdf.SetX(spaceCustomerInfo)
	pdf.Text("Order no ")
	pdf.SetX(spaceTitik)
	pdf.Text(":")
	pdf.SetX(spaceValue)
	pdf.Text(dataRptHdr.OrderNo)

}
