package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"webGinDemo/web09/dao"
	"webGinDemo/web09/model"
	"webGinDemo/web09/routers"
)



func main()  {
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()	// 程序退出关闭数据库连接
	dao.DB.AutoMigrate(&model.Subtodo{})

	router := routers.SetupRouter()
	router.Run()
}