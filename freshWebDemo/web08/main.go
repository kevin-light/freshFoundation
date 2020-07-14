package main

import "github.com/gin-gonic/gin"

type Subtodo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}

func main()  {
	router := gin.Default()
	router.POST("/", func(ctx *gin.Context) {
		var subtodo Subtodo
		if ctx.ShouldBind(&subtodo) == nil {
			ctx.JSON(200,gin.H{"status": "200"})
		}else {
			ctx.JSON(401,gin.H{"status":"400"})
		}

	})
	router.Run()
}