package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	v1 "resto-be/controllers/v0.1.0"
	"resto-be/utils"
)

type RestoBeServiceEnv struct {
	ReadTo  int `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTo int `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	version string

	accessPointTest string
	accessPointResto string

	nameService    string

	debugMode string
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

func InitRouter() *gin.Engine  {
	r := gin.New()

	fmt.Println(gin.IsDebugging())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET(accessPointTest, v1.TestController)


	var api *gin.RouterGroup

	AuthController := new(v1.AuthController)
	api = r.Group(version + "/auth")
	api.POST("/login", AuthController.Login)


	RestoController := new(v1.RestoController)
	api = r.Group(version + "/resto")
	api.POST("/", RestoController.Save)
	api.GET("/all", RestoController.GetAll)
	api.POST("/page/:page/count/:count", RestoController.GetByFilterPaging)

	//r.POST(accessPointResto)

	return r

}