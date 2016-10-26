package main

import (
	_ "ZyyGo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/layui", "static/layui")
	beego.SetStaticPath("/scripts", "static/scripts")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
