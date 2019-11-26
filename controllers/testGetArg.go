package controllers

import (
	"github.com/astaxie/beego"
)

type User struct {
	Id       int    `form:"_"` //忽略id的 还可以小写id
	Username string `form:"username"`
	Age      string `form:"age"`
	Email    string `form:"email"`
}

type TestArgController struct {
	beego.Controller
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
func (c *TestArgController) Get() {
	c.TplName = "index.tpl"

}

func (c *TestArgController) Post() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {

	}
	c.Ctx.WriteString("Email:" + u.Email + "username:" + u.Username + "Age:" + u.Age)

}
