package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"webGinDemo/web09/model"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/


func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"index.html",nil)
}

func CreateSubtodo(ctx *gin.Context) {
	// 获取前端API中的请求数据
	var subtodo model.Subtodo
	ctx.BindJSON(&subtodo)
	// 保存数据库返回响应
	if err := model.CreateSubtodo(&subtodo); err!=nil {
		ctx.JSON(http.StatusOK,gin.H{"error": err.Error()})
	}else {
		ctx.JSON(http.StatusOK,subtodo)
	}
}

func GetSubtodo(ctx *gin.Context) {
	//if DB.Find(&subtodoList) == nil {		// DB.Find 直接返回结果，不能直接==nill
	//	fmt.Println(subtodoList,"---1")
	//	ctx.JSON(http.StatusOK,subtodoList)
	//}else {
	//	fmt.Println(subtodoList,"---2")
	//	ctx.JSON(http.StatusOK,gin.H{"status":"error"})
	//}
	if subtodoList, err := model.GetSubtodoAll(); err != nil {
		fmt.Println(subtodoList,"---1")
		ctx.JSON(http.StatusOK,gin.H{"status":"error"})
	}else {
		ctx.JSON(http.StatusOK,subtodoList)
	}
}

func PutSubtodo(ctx *gin.Context) {
	id,ok := ctx.Params.Get("id")
	if !ok{
		ctx.JSON(http.StatusOK,gin.H{"error":"无效的ID"})
		return
	}
	subtodo,err := model.GetSubtodo(id)
	if err!=nil{
		ctx.JSON(http.StatusOK,gin.H{"err":err.Error()})
		return
	}
	ctx.BindJSON(&subtodo)
	fmt.Println(subtodo,"---3")
	if err = model.UpdateSubtodo(subtodo); err!=nil{
		ctx.JSON(http.StatusOK,gin.H{"err":err.Error()})
		return
	}else {
		ctx.JSON(http.StatusOK,subtodo)
	}

}

func DeleteSubtodo(ctx *gin.Context) {
	id,ok := ctx.Params.Get("id")
	fmt.Println(id,"---2")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{"err":"无效的ID"})
		return
	}
	if err := model.DeleteSubtodo(id); err != nil {
		ctx.JSON(http.StatusOK,gin.H{"err":err.Error()})
	}else {
		ctx.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}