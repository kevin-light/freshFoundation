package routers

import (
	"github.com/gin-gonic/gin"
	"webGinDemo/web09/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// gin框架模板文件和静态文件引用
	router.Static("/static","static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)
	v1Group := router.Group("v1")
	{
		// 待办事项 添加
		v1Group.POST("/todo", controller.CreateSubtodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetSubtodo)
		// 参看一个待办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {
		})
		// 修改代办事项
		v1Group.PUT("/todo/:id", controller.PutSubtodo)
		// 删除一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteSubtodo)
	}
	return router
}