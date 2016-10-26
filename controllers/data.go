package controllers

import (
	"fmt"
	. "zyygo/models"
	_ "zyygo/services"

	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}

func (c *DataController) Menu() {
	var xiangmuname = c.GetString("xiangmuname")
	var xiugai = new(T_xiugai_log).Init()
	var m = xiugai.Query("select * from T_xiugai_log where xiangmuname=" + fmt.Sprint(xiangmuname))
	c.Data["json"] = m
	//c.TplName="main/main.html"
	c.ServeJSON()
}
