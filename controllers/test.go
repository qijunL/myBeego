package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString("这是一个beego控制器Get方法!") //控制器必须对应路由
}

func (c *TestController) Post() {
	c.Ctx.WriteString("这是一个beego控制器Post方法!") //控制器必须对应路由
}