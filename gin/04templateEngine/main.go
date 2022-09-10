package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	//加载模版文件
	engine.LoadHTMLGlob("templates/*")
	//以下这种方式也可以啦
	//engine.LoadHTMLFiles("templates/index.tmpl")
	engine.GET("/index", func(context *gin.Context) {
		//根据文件名渲染
		//最终json将title替换
		context.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "我的标题"})
	})
	engine.Run(":8000")
}
