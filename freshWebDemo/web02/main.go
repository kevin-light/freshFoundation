package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getRequestApi()  {		// get请求
	r := gin.Default()	//Default返回一个默认的路由引擎
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求携带的query setting参数
		//name := c.DefaultQuery("username","默认的返回值")
		name := c.Query("username")
		age := c.Query("age")

		//name,ok := c.GetQuery("username")   //  源代码返回bool，go中bool一般用ok接收
		//if !ok{
		//	// 取不到数据
		//	name = "null"
		//}

		fmt.Println(name,"---0")
		c.JSON(http.StatusOK,gin.H{
			"name" : name,
			"age" : age,
		})
	})
	r.Run(":9000")
}
func postRequestApi() {  		// post请求
	r := gin.Default()	//Default返回一个默认的路由引擎
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单数据 1
		//name := c.PostForm("username")
		//pwd := c.PostForm("password")
		// 获取form表单数据 2
		name := c.DefaultPostForm("username","unname")
		//pwd := c.DefaultPostForm("行行行","....")
		pwd, ok := c.GetPostForm("password")
		if !ok {
			pwd = "sub"
		}

		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name": name,
			"pwd": pwd,
		})
		//	c.JSON(http.StatusOK,gin.H{
		//		"name":name,
		//		"pwd":pwd,
		//})
	})

	r.Run(":9000")
}
func main() {
	r := gin.Default()	//Default返回一个默认的路由引擎
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取途径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})			// /user/:name/:age 和 /blog/:name/:age 长度必须要一样
	r.GET("/blog/:yy/:mm", func(c *gin.Context) {
		// 获取途径参数
		name := c.Param("yy")
		age := c.Param("mm")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})
	r.Run(":9000")
}
