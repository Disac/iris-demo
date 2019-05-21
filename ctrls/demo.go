package ctrls

import (
	"github.com/kataras/iris"
	"go-iris-frame-demo/services"
)

var DemoCtrl = NewDemoCtrl()

func NewDemoCtrl() *Demo {
	return &Demo{
		UserServ: services.UserServ,
	}
}

type Demo struct {
	UserServ *services.User
}

func (ctrl *Demo) Login(ctx iris.Context) {
	loginName := ctx.PostValue("login_name")

	exist, err := ctrl.UserServ.CheckExist(loginName)
	if err != nil {

	}
	if !exist {

	}
}
