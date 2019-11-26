package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //操作mysql
)

//这里数据库中user_info对应UserInfo
var (
	db orm.Ormer //声明衣一个orm对象
)

type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func init() {
	//只做一次操作
	orm.Debug = true //是否开启调试模式,在调试模式下可以打印出SQL语句
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(localhost:3306)/test?parseTime=true")
	orm.RegisterModel(new(UserInfo)) //创建一个表
	db = orm.NewOrm()
}

//添加一个对象
func AddUser(user *UserInfo) (int64, error) {
	id, err := db.Insert(user)
	return id, err
}

//读取一个队象
func ReadUserInfo(users *[]UserInfo) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}
