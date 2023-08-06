package routers

import (
	"github.com/beego/beego/v2/server/web"
	"weblog/controllers"
)

func init() {

	web.Router("/", &controllers.BlogController{}, "*:Home")
	web.Router("/home", &controllers.BlogController{}, "*:Home")
	web.Router("/article", &controllers.BlogController{}, "*:Article")
	web.Router("/detail", &controllers.BlogController{}, "*:Detail")
	web.Router("/about", &controllers.BlogController{}, "*:About")
	web.Router("/timeline", &controllers.BlogController{}, "*:Timeline")
	web.Router("/resource", &controllers.BlogController{}, "*:Resource")
	web.Router("/comment", &controllers.BlogController{}, "post:Comment")
	web.AutoRouter(&controllers.AdminController{})

}
