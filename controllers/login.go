package controllers

import (
"fmt"
"github.com/astaxie/beego"
)

type LoginController struct{
    beego.Controller
}
//登录页面
func (c *LoginController0) Get()  {
    fmt.Println("denglu")
}