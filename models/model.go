package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

//表的设计

//定义一个结构体
type User struct {
	Id int
	Name string
	PassWord string
	//Pass_Word
}

type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string  `orm:"size(100)"`
}

//在ORM里面__是有特殊含义的


func init(){
//ORM操作数据库
//获取连接对象
orm.RegisterDataBase("default","mysql","root:yang1234@tcp(127.0.0.1:3306)/go_db_1?charset=utf8")

//创建表
orm.RegisterModel(new(User))

//生成表
//第一个参数是数据库别名，第二个参数是是否强制更新
orm.RunSyncdb("default",false,true)
//操作表

}