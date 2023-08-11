package main

import (
	"GiveBook/common/dao"
	"GiveBook/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//初始化数据库
	DB := dao.InitDB()
	//关闭数据库
	defer DB.Close()
	r := routes.SetupRoute()
	r.Run(":8080")
}
