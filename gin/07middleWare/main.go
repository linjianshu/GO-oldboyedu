package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//定义全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了...")
		//设置变量到context的key中,可以通过get取
		context.Set("request", "中间件")
		//执行函数
		context.Next()
		//中间件执行完后续的一些事情
		status := context.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		fmt.Println("用时 : ", time.Now().Sub(t))
	}
}
func main() {
	engine := gin.Default()

	//注册中间件
	engine.Use(MiddleWare())
	{
		engine.GET("/middleware", func(context *gin.Context) {
			//取值
			request, _ := context.Get("request")
			fmt.Println(request)
			//页面接收
			context.JSON(http.StatusOK, gin.H{"request": request})
		})

		//如此就是在这个路由之下定义了新的中间件,可以视为单个中间件
		engine.GET("/middlewareDouble", MiddleWare(), func(context *gin.Context) {
			//取值
			request, _ := context.Get("request")
			fmt.Println(request)
			//页面接收
			context.JSON(http.StatusOK, gin.H{"request": request})
		})
	}
	engine.Run(":8000")
}
