package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"go-iris-frame-demo/ctrls"
)

func init() {
	Route.register(&UserRouter{
		app:      Route.app,
		DemoCtrl: ctrls.DemoCtrl,
	})
}

type UserRouter struct {
	app      *iris.Application
	DemoCtrl *ctrls.Demo
}

func (r *UserRouter) Register() {
	r.app.PartyFunc("/user", func(user router.Party) {
		user.Get("/login", r.DemoCtrl.Login)
	})
}

func (r *UserRouter) Name() string {
	return "user"
}
