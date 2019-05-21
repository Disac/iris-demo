package main

import (
	"go-iris-frame-demo/routes"

	_ "go-iris-frame-demo/etc/config"
	_ "go-iris-frame-demo/provider/mysql"
	_ "go-iris-frame-demo/provider/rds"
)

func main() {
	routes.Route.Init()
}
