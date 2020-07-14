package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadFileOne()  {			// 单个文件
	// 单个文件
	router := gin.Default()
	router.LoadHTMLFiles("./index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	// 处理multipart forms 提交文件时默认的内存限制是32Mib，修改方式：router.MaxMultipartMemory=8<<20
	router.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file,err := c.FormFile("fi")
		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		//dst := fmt.Sprintf("C:/WORK/dev_go/src/web04/tmp/%s",file.Filename)
		dst := fmt.Sprintf("./%s",file.Filename)
		// 上传文件到指定目录
		c.SaveUploadedFile(file,dst)
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("'%s' uploaded",file.Filename),
		})
	})
	router.Run(":9000")

}
func main() {		//  多个文件上传 postman提交

	router := gin.Default()
	router.LoadHTMLFiles("./index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	// 处理multipart forms 提交文件时默认的内存限制是32Mib，修改方式：router.MaxMultipartMemory=8<<20
	router.POST("/upload", func(c *gin.Context) {
		// multipart form 多个文件上传
		form, _:= c.MultipartForm()
		files := form.File["file"]
		for index,file := range files{
			log.Println(file.Filename)
			dst:=fmt.Sprintf("./tmp/%s_%d",file.Filename,index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file,dst)
		}
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("'%s' uploaded", len(files)),

		})
	})
	router.Run(":9000")
}
