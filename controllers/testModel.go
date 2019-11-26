package controllers

import (
	"github.com/astaxie/beego"
	"myproject/models"
	//"myproject/models"
)

//这里数据库中user_info对应UserInfo

type TestModelController struct {
	beego.Controller
}

//func (c *TestModelController) Get(){
//	//插入
//	/*
//	newOrm := orm.NewOrm()
//	User := UserInfo{Username:"zhangsan",Password:"123456"}
//	id,err := newOrm.Insert(&User)
//	c.Ctx.WriteString(fmt.Sprintf("操作数据库id=%v",id,err))
//
//	User = UserInfo{Username:"zhangsan",Password:"123456"}
//	id,err = newOrm.Insert(&User)
//	c.Ctx.WriteString(fmt.Sprintf("操作数据库id=%v",id,err))
//	//查询 修改
//	User := UserInfo{Username:"lisi",Password:"123456"}
//	User.Id= 1
//	newOrm := orm.NewOrm()
//	newOrm.Update(&User)
//	c.Ctx.WriteString(fmt.Sprintf("操作数据库user:%v",User)) */
//	//查询
//	newOrm := orm.NewOrm()
//	User := UserInfo{Id:1}
//	newOrm.Read(&User)
//	c.Ctx.WriteString(fmt.Sprintf("操作数据库user:%v",User))
//
//
//}

func (c *TestModelController) Get() {
	//user := models.UserInfo{Username:"王五",Password:"123456",Id:2}
	//models.AddUser(&user)
	//c.Ctx.WriteString("call model success")
	var users []models.UserInfo
	models.ReadUserInfo(&users)
	c.Data["Title"] = "zhangsan"
	c.Data["Users"] = users
	c.Data["len"] = len(users)
	c.TplName = "test.tpl"
}
