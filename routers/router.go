package routers

import (
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Test")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test/arg", &controllers.TestArgController{}, "get:Get;post:Post")
	//orm
	beego.Router("/test/orm", &controllers.TestModelController{}, "get:Get;post:Post")
}
