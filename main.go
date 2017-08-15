package main

import (
	_ "DDFashion/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:zhaoyi@/dd_fashion?charset=utf8")
	orm.RunSyncdb("default", true, true)
	orm.Debug = beego.BConfig.RunMode == "dev"
}
