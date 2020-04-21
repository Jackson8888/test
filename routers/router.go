package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 新增路由
	beego.Router("/hello", &controllers.HelloController{})
}
