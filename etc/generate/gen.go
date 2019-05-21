package main

import (
	"go-iris-frame-demo/etc/generate/gen"
)

func main() {
	gen.NewGen().
		SpecialTables("user").UseInterface().
		G("go-frame-demo/models", "root:123456@tcp(localhost:3306)/pay?parseTime=true&loc=Local")
}
