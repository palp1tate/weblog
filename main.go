package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "weblog/models"
	_ "weblog/routers"
)

func main() {
	web.Run()
}
