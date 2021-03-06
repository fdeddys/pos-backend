package routers

import (
	"fmt"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/utils"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"

	v1 "resto-be/controllers/v0.1.0"
)

type RestoBeServiceEnv struct {
	ReadTo  int `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTo int `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	version string

	accessPointTest  string
	accessPointResto string

	nameService string

	debugMode         string
	restoBeServiceEnv RestoBeServiceEnv
)

func init() {
	version = "/v0.1.0"

	accessPointTest = version + "/test"

	debugMode = utils.GetEnv("APPS_DEBUG", "debug")

	err := envconfig.Process("RESTO_BE_SERVICE", &restoBeServiceEnv)
	if err != nil {
		fmt.Println("Failed to get RESTO_BE_SERVICE env:", err)
	}

	nameService = utils.GetEnv("RESTO_BE_SERVICE", "rose-be-service")
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	fmt.Println(gin.IsDebugging())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET(accessPointTest, v1.TestController)

	var api *gin.RouterGroup

	AuthController := new(v1.AuthController)
	api = r.Group(version + "/auth")
	api.POST("/login", AuthController.Login)
	api.GET("/:id/reset-password", cekToken, AuthController.ResetPassword)

	RestoController := new(v1.RestoController)
	api = r.Group(version + "/resto")
	api.POST("", cekToken, RestoController.Save)
	api.POST("/savebyresto", cekToken, RestoController.SaveByResto)
	api.POST("/getselfresto", cekToken, RestoController.GetByRestoToken)
	api.GET("", cekToken, RestoController.GetAll)
	api.GET("/:id", RestoController.GetById)
	api.POST("/check-code", RestoController.CheckCodeResto)
	api.POST("/upload-image", cekToken, RestoController.UploadImage)
	api.POST("/remove-image", cekToken, RestoController.RemoveImage)
	api.POST("/page/:page/count/:count", RestoController.GetByFilterPaging)

	//TaxController := new(v1.TaxController)
	//api = r.Group(version + "/tax")

	EMenuGroupController := new(v1.EMenuGroupController)
	api = r.Group(version + "/menu-group")
	api.POST("", cekToken, EMenuGroupController.Save)
	api.GET("", EMenuGroupController.GetAll)
	api.POST("/upload-image", cekToken, EMenuGroupController.UploadImage)

	api.GET("/id/:id", EMenuGroupController.GetById)
	api.POST("/resto/:restoId/page/:page/count/:count", EMenuGroupController.GetByFilterPaging)
	api.GET("/resto/:restoId", EMenuGroupController.GetByIdResto)
	api.POST("/filter", EMenuGroupController.Filter)

	CategoryController := new(v1.CategoryController)
	api = r.Group(version + "/category")
	api.GET("", CategoryController.GetAll)

	EMenuItemController := new(v1.EMenuItemController)
	api = r.Group(version + "/menu-item")
	api.POST("", cekToken, EMenuItemController.Save)
	api.GET("", EMenuItemController.GetAll)
	api.GET("/id/:id", EMenuItemController.GetById)
	api.GET("/menu-group/:groupId", EMenuItemController.GetByMenuGroupId)
	api.GET("/resto/:restoId", EMenuItemController.GetByRestoId)
	api.GET("/favorite/resto/:restoId", EMenuItemController.GetFavoriteByRestoId)
	api.POST("/page/:page/count/:count", EMenuItemController.GetByFilterPaging)
	api.POST("/upload-image", cekToken, EMenuItemController.UploadImage)
	api.POST("/remove-image", cekToken, EMenuItemController.RemoveImage)
	api.POST("/filter", EMenuItemController.Filter)

	//api.GET("/menu-group/:groupId/resto/:restoId", EMenuItemController.GetByMenuGroupIdAndIdResto)

	VoucherController := new(v1.VoucherController)
	api = r.Group(version + "/voucher")
	api.POST("", cekToken, VoucherController.Save)
	api.POST("/getbycode", VoucherController.GetByCodeVoucher)
	api.GET("/:id", VoucherController.GetById)
	api.POST("/page/:page/count/:count", cekToken, VoucherController.GetByFilterPaging)

	OrderController := new(v1.OrderController)
	api = r.Group(version + "/order")
	api.POST("/add", OrderController.Add)
	api.POST("/customer/page/:page/count/:count", OrderController.GetByCustPage)
	//api.POST("/resto/:restoId/page/:page/count/:count", OrderController.GetByRestoFilterPaging)
	api.GET("/preview/:id", OrderController.PrintPreview)
	api.GET("/id/:id", OrderController.GetByID)
	api.GET("/detail/:orderId", OrderController.GetOrderDetailByOrderId)
	api.POST("/page/:page/count/:count", OrderController.GetByFilterPaging)

	// ADD NEW
	// 1. get order (preload=> order-detail) by tabel no
	api.GET("/restoId/:restoId/tabel/:tabelId", OrderController.GetByRestoIdTabelID)
	// 3. add item by no tabel
	api.POST("/addItem", OrderController.AddItemOrderToTabel)
	// 5. payment by tabel no ( cash + debit + cc)
	api.POST("/payment/:tabelID", OrderController.PaymentByTabelID)
	api.GET("/payment/:tabelID", OrderController.GetPaymentByTabelID)

	api.POST("/update-status", OrderController.UpdateStatusOrder)
	api.POST("/update-status-complete", OrderController.UpdateStatusCompleteOrder)
	api.POST("/update-status-detail", OrderController.UpdateStatusOrderDetail)
	api.POST("/update-qty", OrderController.UpdateQty)
	UserController := new(v1.UserController)
	api = r.Group(version + "/user")
	api.POST("/page/:page/count/:count", cekToken, UserController.GetByFilterPaging)
	api.POST("/", cekToken, UserController.SaveUser)
	api.GET("/reset-password/:id", cekToken, AuthController.ResetPassword)
	api.POST("/change-password", cekToken, UserController.ChangePassword)

	api.GET("/current-user", cekToken, UserController.GetUser)

	//api.GET("/", EMenuItemController.GetAll)
	//api.GET("/:id", EMenuItemController.GetById)
	//api.POST("/page/:page/count/:count", EMenuItemController.GetByFilterPaging)

	//r.POST(accessPointResto)

	ImageController := new(v1.ImageController)
	api = r.Group(version + "/image")
	api.POST("/upload", cekToken, ImageController.Upload)

	CustomerController := new(v1.CustomerController)
	api = r.Group(version + "/customer")
	api.POST("/login", CustomerController.Login)
	api.POST("", CustomerController.Save)
	api.POST("/savebyresto", cekToken, CustomerController.Save)
	api.POST("/page/:page/count/:count", cekToken, CustomerController.FilterPage)
	api.GET("/id/:id", CustomerController.GetByID)

	ReportController := new(v1.ReportController)
	api = r.Group(version + "/report")
	api.POST("/order", cekToken, ReportController.Order)
	api.POST("/order/page/:page/count/:count", cekToken, ReportController.GetOrderByFilterPaging)
	api.POST("/order/detail", cekToken, ReportController.OrderDetail)

	//VoucherController := new(v1.VoucherController)
	//api = r.Group(version + "/voucher")
	//api.POST("", cekToken, VoucherController.Save)
	//api.POST("/getbycode",VoucherController.GetByCodeVoucher)
	//api.GET("/:id", VoucherController.GetById)
	//api.POST("/page/:page/count/:count", cekToken, VoucherController.GetByFilterPaging)

	GroupTableController := new(v1.GroupTableController)
	api = r.Group(version + "/group-table")
	api.POST("", cekToken, GroupTableController.Save)
	//api.POST("/page/:page/count/:count", cekToken, GroupTableController.GetByFilterPaging)
	api.POST("/filter", GroupTableController.Filter)

	TableController := new(v1.TableController)
	api = r.Group(version + "/table")
	api.POST("", cekToken, TableController.Save)
	api.POST("/filter", TableController.Filter)

	PaymentTypeCtrl := new(v1.PaymentTypeController)
	api = r.Group(version + "/payment-type")
	api.POST("/get-all", PaymentTypeCtrl.GetAll)

	// ADD NEW
	// 1. get order (preload=> order-detail) by tabel no
	// 2. get all active table ( status table open/occupacy jgn lupa ) => sdh ada
	// 3. add item by no tabel
	// 4. update qty item by table no/ order detail id - no add/substrac qty
	// 5. payment by tabel no ( cash + debit + cc)

	// 6. join tabel & separate tabel
	// 7. move tabel
	// 8. print bill, caption & cur bill ( bill w/o payment)
	// 9. print report
	// 10.login cashier
	// 11.add waiters

	//api.POST("/customer/page/:page/count/:count", OrderController.GetByCustPage)

	return r

}

func cekToken(c *gin.Context) {

	res := models.Response{}
	tokenString := c.Request.Header.Get("Authorization")
	log.Println("tokenString -> ", tokenString)

	if strings.HasPrefix(tokenString, "Bearer ") == false {
		res.Rc = constants.ERR_CODE_53
		res.Msg = constants.ERR_CODE_53_MSG + " [01] "
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if jwt.GetSigningMethod("HS256") != token.Method {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [02] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			// return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.TokenSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("claims : ", claims)

		fmt.Println("User name from TOKEN ", claims["user"])

		unixNano := time.Now().UnixNano()
		timeNowInInt := unixNano / 1000000

		tokenCreated := (claims["tokenCreated"])
		dto.CurrUserEmail = (claims["userEmail"]).(string)

		currUserId := (claims["userId"]).(string)
		dto.CurrUserID, _ = strconv.ParseInt(currUserId, 10, 64)

		currRestoId := (claims["restoId"]).(string)
		dto.CurrRestoID, _ = strconv.ParseInt(currRestoId, 10, 64)

		fmt.Println("now : ", timeNowInInt)
		fmt.Println("token created time : ", tokenCreated)
		fmt.Println("user by token : ", dto.CurrUserEmail)
		fmt.Println("user by token ID : ", dto.CurrUserID)

		tokenCreatedInString := tokenCreated.(string)
		tokenCreatedInInt, errTokenExpired := strconv.ParseInt(tokenCreatedInString, 10, 64)

		if errTokenExpired != nil {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [03] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if ((timeNowInInt - tokenCreatedInInt) / 1000) > constants.TokenExpiredInMinutes {
			res.Rc = constants.ERR_CODE_53
			res.Msg = constants.ERR_CODE_53_MSG + " [04] "
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		fmt.Println("Token already used for ", (timeNowInInt-tokenCreatedInInt)/1000, "sec, Max expired ", constants.TokenExpiredInMinutes, "sec ")
		// fmt.Println("token Valid ")

	} else {
		res.Rc = constants.ERR_CODE_53
		res.Msg = constants.ERR_CODE_53_MSG + " [05] "
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
}
