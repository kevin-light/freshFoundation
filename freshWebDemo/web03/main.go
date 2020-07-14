package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserInfo struct {
	Username string `form:"username" json:"user"`  // form
	Password string `form:"password" json:"pwd"`
}
func demo1()  {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo			// 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo			// 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.POST("/json", func(c *gin.Context) {		// json格式
		var u UserInfo			// 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"json",
			})
		}
	})
	r.Run(":9000")
}

type Loginuser struct {
	User     string `form:"user" json:"user" binding:"required"`       //  tag 加上了 binding:"required"，但绑定时是空值, Gin 会报错。
	Password string `form:"password" json:"pwd" binding:"required"`
}

func main()  {
	router := gin.Default()
	// 绑定json的示例({"user": "qqq", "password": "123456"})
	router.POST("/loginJson", func(c *gin.Context) {
		var login Loginuser
		if err := c.ShouldBind(&login); err == nil{
				fmt.Printf("login info:%#v\n",login)
				c.JSON(http.StatusOK,gin.H{
					"user":login.User,
					"pwd":login.Password,
				})
		}else {
			c.JSON(http.StatusBadRequest,gin.H{"err:":err.Error()})
		}
	})
	// 绑定form 表单 (user=lll&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Loginuser
		if err:= c.ShouldBind(&login);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err:":err.Error(),
			})
			return
		}
		if login.User != "qqq" || login.Password != "123123"{
			c.JSON(http.StatusUnauthorized,gin.H{"status":"unauthorized"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"you are login in "})
	})
	router.Run(":9000")

}
