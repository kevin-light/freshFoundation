package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func rqt1()  {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET:请求方式；/hello ：请求路径
	// 当客户端以GET方式请求hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON : 返回json格式的数据
		c.JSON(200,gin.H{
			"message": "hello world",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
func rqt2()  {
	// gin框架 restfulApi的开发
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"get",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"post",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"put",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"delete",
		})
	})
	r.Run()
}

//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
func rqt3()  {
	rqt := gin.Default()
	rqt.Static("/xxx","./statics")	// 加载静态文件

	rqt.LoadHTMLGlob("templates/**/*")		// **代表文件夹, *代表文件夹下的所有文件
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	rqt.GET("/posts/index", func(ct *gin.Context) {
		ct.HTML(http.StatusOK,"postsss/index.html",gin.H{
			"title":" posts -> index ",
		})
	})
	rqt.GET("users/index", func(ct *gin.Context) {
		ct.HTML(http.StatusOK,"usersss/index.html",gin.H{
			"title":"users -> index",
		})
	})
	// 返回下载的前端模板
	rqt.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"home.html",nil)
	})
	rqt.Run(":9000")

}

//自定义模板函数 :  定义一个不转义相应内容的safe模板函数如下


func rqt4()  {
	router := gin.Default()
	router.Static("/xxx","./statics")	// 加载静态文件

	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html","<a href='https://liwenzhou.com'>GO博客</a>")
	})
	router.Run(":8000")
}

//静态文件处理 : 当我们渲染的HTML文件中引用了静态文件时，我们只需要按照以下方式在渲染页面前调用gin.Static方法即可。
func rqt5()  {
	r := gin.Default()
	r.Static("/static","./statics")
	r.LoadHTMLGlob("templates/**/*")
}
func main()  {
	//rqt1()
	//rqt2()
	rqt3()
	//rqt4()
}
