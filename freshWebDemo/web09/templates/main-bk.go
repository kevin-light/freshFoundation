package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Subtodo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}

var (
	DB *gorm.DB
)

func initMysql() (err error) {
	dsn := "root:123123@tcp(127.0.0.1:3306)/dtdba?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main()  {
	// 连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()	// 程序退出关闭数据库连接
	DB.AutoMigrate(&Subtodo{})

	router := gin.Default()
	// gin框架模板文件和静态文件引用
	router.Static("/static","static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"index.html",nil)
	})
	v1Group := router.Group("v1")
	{
		// 待办事项 添加
		v1Group.POST("/subtodo", func(ctx *gin.Context) {
			// 获取前端API中的请求数据
			var subto Subtodo
			ctx.BindJSON(&subto)
			// 保存数据库返回响应
			if err = DB.Create(&subto).Error; err != nil {
				ctx.JSON(http.StatusOK,gin.H{"error": err.Error()})
			}else {
				ctx.JSON(http.StatusOK,subto)
			}
		})

		// 查看所有的待办事项
		v1Group.GET("/subtodo", func(ctx *gin.Context) {
			//var subtodoList []Subtodo
		})
		// 参看一个待办事项
		v1Group.GET("/subtodo/:id", func(ctx *gin.Context) {

		})
		// 修改代办事项
		v1Group.PUT("/subtodo/:id", func(ctx *gin.Context) {

		})
		// 删除一个代办事项
		v1Group.DELETE("/subtodo/:id", func(ctx *gin.Context) {

		})
	}


	router.Run()
}