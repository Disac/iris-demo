package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

type Router interface {
	Register()
	Name() string
}

var Route = NewRoute()

func NewRoute() *route {
	app := iris.New()
	app.Use(recover.New())

	router := &route{
		app:     app,
		handler: make(map[string]Router),
	}
	return router
}

type route struct {
	app     *iris.Application
	handler map[string]Router
}

func (r *route) Init() {
	for _, router := range r.handler {
		router.Register()
	}

	err := r.app.Run(iris.Addr(":8080"),
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithCharset("UTF-8"),
	)
	if err != nil {
		panic(err)
	}
}

func (r *route) register(router Router) {
	r.handler[router.Name()] = router
}
