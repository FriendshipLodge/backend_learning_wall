package controllers

import (
	"github.com/astaxie/beego"
	"resbox_go/models"
)

type UserInfoController struct {
	beego.Controller
}

func (u *UserInfoController) Get() {
	u.Data["json"]=models.ResponseMod{
		Code:    "401",
		Message: "请使用正确的请求方式",
		Data:    nil,
	}
	u.ServeJSON()
}

func (u *UserInfoController) Post()  {
	UserList,_ := models.GetUserInfoById(2)
	u.Data["json"] = models.ResponseMod{
		Code:    "200",
		Message: "查询成功",
		Data:    UserList,
	}
	u.ServeJSON()
}

func (u *UserInfoController) GetAll() {
	UserList,_ := models.GetAllUserInfo()
	u.Data["json"] = models.ResponseMod{
		Code:    "200",
		Message: "查询成功",
		Data:    UserList,
	}
	u.ServeJSON()
}