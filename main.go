package main

import (
	"fmt"
	"log"
	"resto-be/routers"
	"resto-be/utils"
	"runtime"
	"strconv"
)

var (
	port string
)

func main()  {

	maxProc, _ := strconv.Atoi(utils.GetEnv("MAXPROCS", "1"))
	port = utils.GetEnv("PORT", "8001")
	runtime.GOMAXPROCS(maxProc)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", port)

	log.Println("[info] start http server listening %s", endPoint)

	//server.ListenAndServe()

	routersInit.Run(":" + port)
}