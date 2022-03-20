package controllers

import (
	"github.com/astaxie/beego"
	"resbox_go/models"
)

type UserInfoController struct {
	beego.Controller
}

func (u *UserInfoController) Healthy() {
	u.Data["json"]=models.ResponseMod{
		Code:    "200",
		Message: "Healthy",
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