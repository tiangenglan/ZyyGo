package controllers

import "github.com/astaxie/beego"
import "fmt"

//UserController 用户控制器
type UserController struct {
	beego.Controller
}

// Gets  表单提交
func (c *UserController) Gets() {
	c.Data["this"] = "beego"
	fmt.Println("kaishile")
}
