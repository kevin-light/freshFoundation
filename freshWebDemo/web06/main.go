package main
//  Gin中的中间件必须是一个gin.HandlerFunc类型。例如我们像下面的代码一样定义一个统计请求耗时的中间件。
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message":"hello world",
	})
}

// // StatCost 是一个统计耗时请求耗时的中间件
func m1(c *gin.Context)  {
	fmt.Println("m1 in ...")
	//return func(c *gin.Context) {
		start := time.Now()
		c.Set("name","hello")  // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Next()		// 调用该请求的剩余处理程序
		//c.Abort()		// 不调用该请求的剩余处理程序
		cost := time.Since(start)		// 计算耗时

	log.Println(cost,"---m1")
	fmt.Printf("cost:%v\n",cost)
	//}
}
func m2(c *gin.Context)  {
	fmt.Println("m2 in ...")
	name,ok:=c.Get("name")  // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	//go funcXxx(c.Copy())  //当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。

	if !ok{
		name="null"
	}
	c.Next()		// 调用该请求的剩余处理程序
	//c.Abort()		// 不调用该请求的剩余处理程序
	//return			// 直接返回后面不执行
	log.Println("---m2",name)
	//}
}
func authMiddleware(doCheck bool) gin.HandlerFunc {		// authMiddleware函数传入bool做控制 true做检查，else部不检查
	// 连接数据库 ,或者其他功能
	return func(c *gin.Context) {
		if doCheck{
			// 存放具体逻辑，
			//判断用户是否登录：=
			//if登录
			//c.Next()
			//else
			//c.About()
		}else {
			c.Next()
		}
	}
}
func main() {
	r:=gin.Default()
	r.Use(m1,m2,authMiddleware(true))		// 全局注册 m1 中间件，不用每个路由都写了
	r.GET("/index", indexHandler)
	//r.GET("/index", m1, indexHandler)

	//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。

	// 路由组的中间件：方法一
	shopGroup := r.Group("/shop1", authMiddleware(false))
	{
		shopGroup.GET("/index1", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"status":"ok"})
		})
	}
	// 路由组的中间件：方法二
	shopGroup2 := r.Group("/shop2", authMiddleware(false))
	shopGroup2.Use(authMiddleware(false))
	{
		shopGroup2.GET("/index2", func(c *gin.Context) {
			shopGroup2.Use(authMiddleware(false))
		})
	}

	r.Run(":9000")
}


// 中间件注意事项::
//gin默认中间件
//gin.Default()默认使用了Logger和Recovery中间件，其中：
//Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
//如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
//gin中间件中使用goroutine -->
//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。



/*
//我们可以在多个端口启动服务，例如：
var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func mainDemo1() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// 借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
*/
