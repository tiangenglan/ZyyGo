package controllers

import (
	"fmt"
	. "zyygo/models"
	//"zyygo/services"
	"github.com/astaxie/beego"
)

//MainController 入库控制器
type MainController struct {
	beego.Controller
}

//Get 主程序入库
func (c *MainController) Gets() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//c.TplName = "main.html"
}

//显示主界面
func (c *MainController) Get() {

	c.TplName = "main/main.html"

}

// 提交表单
func (c *MainController) Post() {
	var xiangmuname = c.GetString("xiangmuname")
	var time = c.GetString("time")
	var xiugairen = c.GetString("xiugairen")
	var xiugaixinxi = c.GetString("xiugaixinxi")
	fmt.Println(xiangmuname, time, xiugairen, xiugaixinxi)
	//c.TplName = "main.html
	var xiugai = new(T_xiugai_log).Init()
	xiugai.Xiangmuname = xiangmuname
	xiugai.Time = time
	xiugai.Xiugainame = xiugairen
	xiugai.Beizhuinfor = xiugaixinxi

	fmt.Println("需要插入的信息为：", xiugai)
	var b = xiugai.Insert()
	if b > 0 {
		//c.Ctx.WriteString("alert('保存成功！');window.location='/main/main.html';")
		c.Ctx.WriteString("alert('保存成功！')")
		//c.Redirect("main/main.html",302)
	} else {
		c.Ctx.WriteString("alert('保存失败！');")
	}
	c.TplName = "main/main.html"
}

//生成菜单列表
func (c *MainController) Menu() {
	var xiangmuname = c.GetString("xiangmuname")
	var xiugai = new(T_xiugai_log).Init()
	var m = xiugai.Query("select * from T_xiugai_log where xiangmuname=" + fmt.Sprint(xiangmuname))
	c.Data["list"] = m
	c.TplName = "main/main.html"
}
