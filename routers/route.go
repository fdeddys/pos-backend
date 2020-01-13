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
	header string

	accessPointTest string
	accessPointLogin string

	nameService    string

	debugMode string
	restoBeServiceEnv RestoBeServiceEnv
)


func init() {
	header = "/v0.1.0"

	accessPointTest = header + "/test"

	accessPointLogin = header + "/login"

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
	r.POST(accessPointLogin, v1.LoginController)


	return r

}