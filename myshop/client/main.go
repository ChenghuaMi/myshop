package main

import (
	_ "myshop/client/bootstrap"
	"myshop/client/core/route"
)

func main() {
	route.InitRoute().Run()
}
