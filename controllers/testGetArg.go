package controllers

import (
	"github.com/astaxie/beego"
)

type TestArgController struct {
	beego.Controller
}
type User struct {
	UserName string
	age int
	Email string
}
//获取参数
/*
func (c *TestSArgController) Get(){
	id := c.GetString("id")
	c.Ctx.WriteString(id)
	name := c.Input().Get("name")
	c.Ctx.WriteString(name)
}
 */
func (c *TestArgController) Get(){
	c.TplName = "index.tpl"

}

func (c *TestArgController) Post(){
	c.TplName = "index.tpl"

}

