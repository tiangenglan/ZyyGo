package routers

import (
	"ZyyGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/main", &controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})
}
