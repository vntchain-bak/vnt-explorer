package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/vntchain/vnt-explorer/common"
	"github.com/vntchain/vnt-explorer/models"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Post() {
	account := &models.Account{}
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, account)
	if err != nil {
		this.ReturnErrorMsg("Wrong format of Account: %s", err.Error())
		return
	}

	err = account.Insert()
	if err != nil {
		this.ReturnErrorMsg("Failed to create Account: %s", err.Error())
	} else {
		this.ReturnData(account)
	}
}

func (this *AccountController) getCond() (int, int) {
	isContract, err := this.GetInt("isContract")
	if err != nil {
		beego.Warn("Failed to read isContract: ", err.Error())
		isContract = -1
	}

	isToken, err := this.GetInt("isToken")
	if err != nil {
		beego.Warn("Failed to read isToken: ", err.Error())
		isToken = -1
	}
	return isContract, isToken
}

func (this *AccountController) List() {
	offset, err := this.GetInt("offset")
	if err != nil {
		beego.Warn("Failed to read offset: ", err.Error())
		offset = common.DefaultOffset
	}

	limit, err := this.GetInt("limit")
	if err != nil {
		beego.Warn("Failed to read limit: ", err.Error())
		limit = common.DefaultPageSize
	}

	isContract, isToken := this.getCond()

	order := this.GetString("order")
	fields := this.getFields()

	account := &models.Account{}
	accounts, err := account.List(isContract, isToken, order, offset, limit, fields)
	if err != nil {
		this.ReturnErrorMsg("Failed to list accounts: %s", err.Error())
	} else {
		this.ReturnData(accounts)
	}

}

func (this *AccountController) Get() {
	//beego.Info("params", this.Ctx.Input.Params())
	address := this.Ctx.Input.Param(":address")
	if len(address) == 0 {
		this.ReturnErrorMsg("Failed to get address", "")
		return
	}
	// TODO 这边暂时未对fields进行处理
	//fields := this.getFields()

	account := &models.Account{}
	dbaccount, err := account.Get(address)
	if err != nil {
		this.ReturnErrorMsg("Failed to read account: %s", err.Error())
	} else {
		this.ReturnData(dbaccount)
	}
}

func (this *AccountController) Count() {
	isContract, isToken := this.getCond()

	account := &models.Account{}
	count, err := account.Count(isContract, isToken)
	if err != nil {
		this.ReturnErrorMsg("Failed to get accounts count: %s", err.Error())
	} else {
		this.ReturnData(count)
	}
}