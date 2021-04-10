package controllers

import (
	"errors"
	"strings"

	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/mengjunwei/go_api_project/logger"
	"github.com/mengjunwei/go_api_project/models"
)

const (
	MaxRequestBodyLength = 1024 * 64
	TokenStr             = "tokenId"
	ErrNoRowsString      = "请求出错，请确认资源存在"
	ErrNoTokenId         = "请提供参数tokenId"
	ErrBodyLimit = "请求体不得大于64M"
)

type JsonRet struct {
	Code   int         `json:"code"`
	Result bool        `json:"result"`
	Err    string      `json:"error"`
	Data   interface{} `json:"data"`
}

type ApiJsonBaseController struct {
	beego.Controller
	result    JsonRet
	loginUser *models.LoginUser
	RecordLog bool
}

func (bc *ApiJsonBaseController) Prepare() {
	//tokenId := bc.Ctx.Input.Query(TokenStr)
	//
	//if tokenId == "" {
	//	bc.SetError(errors.New(ErrNoTokenId))
	//	bc.Finish()
	//	return
	//}
}

func (bc *ApiJsonBaseController) Finish() {
	bc.Data["json"] = bc.result
	bc.ServeJSON()
}

func (bc *ApiJsonBaseController) SetErrorWithData(err string, data interface{}) {
	if strings.Contains(err, orm.ErrNoRows.Error()) {
		err = ErrNoRowsString
	}

	bc.result.Err = err
	bc.result.Result = false
	bc.result.Data = data
	bc.result.Code = 0
}

func (bc *ApiJsonBaseController) SetError(err error) {
	errStr := err.Error()
	if strings.Contains(errStr, orm.ErrNoRows.Error()) {
		errStr = ErrNoRowsString
	}

	bc.result.Err = errStr
	bc.result.Result = false
	bc.result.Code = 0

	bc.Finish()
}

func (bc *ApiJsonBaseController) SetData(data interface{}) {
	bc.result.Data = data
	bc.result.Result = true
	bc.result.Err = ""
	bc.result.Code = 1

	bc.Finish()
}

func (bc *ApiJsonBaseController) CheckErr(err error) bool {
	if err != nil {
		logger.Debug(err.Error())
		return false
	} else {
		return true
	}
}

func (bc *ApiJsonBaseController) GetPageInfo() (page int, size int, withCount bool, err bool) {
	page, _ = bc.GetInt("page", 0)
	size, _ = bc.GetInt("pageSize", 10)
	if page < 1 || size < 1 {
		bc.SetError(errors.New("page 或者 pageSize 传参错误"))
		err = true
	}
	withCount, _ = bc.GetBool("withCount", false)
	return
}
