package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpRender()  {    //HTTP 重定向很容易。 内部、外部重定向均支持。

	r:=gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	r.Run(":9000")
}
func urlRender()  {		// 路由重定向，使用HandleContext：

	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		// 指定重定向的url; url地址不变，页面html改变
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"hello":"world"})
	})
	r.Run(":9000")
}
func main() {
	// 将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。

	r:=gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
		// 嵌套路由组
		xx := userGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
	}

	r.Run(":9000")
}
