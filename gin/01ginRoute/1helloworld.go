package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//1.创建路由
	//默认使用了2个中间件 Logger(), Recovery()
	engine := gin.Default()
	//engine:=gin.New()也可以的
	//2.绑定路由规则,执行的函数
	//gin.context , 封装了request和response

	//路由组 实际上就是便于管理 少写点东西
	routerGroup := engine.Group("/Get")
	routerGroup.GET("/v1", func(context *gin.Context) {
		query := context.DefaultQuery("name", "getParam")
		context.String(http.StatusOK, query)

	})

	group := engine.Group("/Post")
	group.POST("/v1", func(context *gin.Context) {
		postParam := context.DefaultPostForm("name", "postParam")
		context.String(http.StatusOK, postParam)
	})

	//index 界面
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!!!")
	})

	//api参数 用 : 来取
	engine.GET("/user/:name/*action", func(context *gin.Context) {
		//
		name := context.Param("name")
		action := context.Param("action")
		context.String(http.StatusOK, name+" is "+action)
	})

	//url参数 ?name="xxx"
	engine.GET("/welcome", func(context *gin.Context) {
		query := context.DefaultQuery("name", "Jack")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s !", query))
	})

	//form表单传参
	engine.POST("/PostForm", PostFormParams)
	//from上传单个文件
	engine.POST("/Upload", UploadFile)
	//限制表单上传大小 8mb,默认值为32mb
	engine.MaxMultipartMemory = 8 << 20
	//form上传多个文件
	engine.POST("UploadFiles", UploadFiles)
	engine.PUT("/xxxput")
	//3.监听端口 默认在8080端口
	engine.Run(":8000")
}

func UploadFiles(context *gin.Context) {
	multipartForm, err := context.MultipartForm()
	if err != nil {
		fmt.Println("received multiple files failed , err :", err)
		context.String(http.StatusBadRequest, fmt.Sprintf("get err %s ", err.Error()))
		return
	}

	//获取所有文件
	files := multipartForm.File["files"]
	//遍历所有files
	for _, file := range files {
		//逐个存
		err := context.SaveUploadedFile(file, file.Filename)
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s ", err.Error()))
			return
		}
	}

	context.String(200, fmt.Sprintf("upload ok %d files!", len(files)))
}

func UploadFile(context *gin.Context) {
	//从表单中取文件
	file, err := context.FormFile("file")
	if err != nil {
		fmt.Println("receive file error, cause : ", err)
		return
	}
	log.Println(file.Filename)
	//传到项目的根目录, 名字就用本身的就好
	err = context.SaveUploadedFile(file, file.Filename)
	if err != nil {
		fmt.Println("file save failed , error : ", err)
		return
	}

	//打印信息
	context.String(200, fmt.Sprintf("'%s' has already uploaded!", file.Filename))
}

func PostFormParams(context *gin.Context) {
	//表单参数 设置默认值
	type1 := context.DefaultPostForm("type", "alert")
	//接收其他的
	userName := context.PostForm("username")
	password := context.PostForm("password")
	//多选框
	hobbys := context.PostFormArray("hobby")

	context.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s , hobbys is %v", type1, userName, password, hobbys))
}
