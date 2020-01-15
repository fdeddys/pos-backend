package v0_1_0

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"resto-be/constants"
	"resto-be/database/dbmodels"
	"resto-be/models"
	"resto-be/models/dto"
	"resto-be/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (controller *UserController) GetByFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> User Controoler - Get All <<<")
	parent := context.Background()
	defer parent.Done()

	req := dto.UserRequesDto{}
	res := models.Response{}

	page, errPage := strconv.Atoi(ctx.Param("page"))
	if errPage != nil {
		log.Println("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	count, errCount := strconv.Atoi(ctx.Param("count"))
	if errCount != nil {
		logs.Info("error", errPage)
		res.Rc = constants.ERR_CODE_02
		res.Msg = constants.ERR_CODE_02_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = services.InitializeUserServiceInterface().GetDataUserByFilterPaging(req, page, count)

	ctx.JSON(http.StatusOK, res)

}

// SaveDataUser ...
func (h *UserController) SaveUser(ctx *gin.Context) {

	fmt.Println(">>> UserController - Save <<<")
	parent := context.Background()
	defer parent.Done()

	req := dbmodels.User{}
	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.Rc = constants.ERR_CODE_03
		res.Msg = constants.ERR_CODE_03_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res = services.InitializeUserServiceInterface().SaveDataUser(&req)

	ctx.JSON(http.StatusOK, res)

	// fmt.Println("entering the save user ")
	// req := dbmodels.User{}
	// res := models.Response{}

	// body := c.Request.Body
	// dataBodyReq, _ := ioutil.ReadAll(body)

	// if err := json.Unmarshal(dataBodyReq, &req); err != nil {
	// 	fmt.Println("Error, body Request ")
	// 	res.ErrCode = constants.ERR_CODE_03
	// 	res.ErrDesc = constants.ERR_CODE_03_MSG
	// 	c.JSON(http.StatusBadRequest, res)
	// 	c.Abort()
	// 	return
	// }

	// c.JSON(http.StatusOK, UserService.SaveUser(&req))
}
