package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	//binding:"required" 修饰的字段,若接收为空值 则报错 是必须字段
	User     string `json:"user" uri:"user" xml:"user" form:"username" binding:"required"`
	Password string `json:"password" uri:"password" xml:"password" form:"password" binding:"required"`
}

func main() {
	//1.创建路由
	engine := gin.Default()
	//2.json绑定
	engine.GET("loginJSON", func(context *gin.Context) {
		//声明接收的变量
		var json Login
		//将request的body中的数据,自动按照json格式解析到结构体
		err := context.ShouldBindJSON(&json)
		if err != nil {
			//返回错误信息
			//gin.H 封装了生成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
			return
		}

		//判断用户名密码是否正确
		if json.User != "2020170281" || json.Password != "lalala123" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	engine.GET("loginForm", func(context *gin.Context) {
		var form Login
		//bind()默认解析并绑定form格式
		//根据请求头中的content-type自动推断
		err := context.Bind(&form)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "2020170281" || form.Password != "lalala123" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	engine.GET("loginURI/:user/:password", func(context *gin.Context) {
		var URI Login
		err := context.ShouldBindUri(&URI)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if URI.User != "2020170281" || URI.Password != "lalala123" {
			context.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": 200})
	})
	engine.Run(":8000")
}
