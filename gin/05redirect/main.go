package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/redirect", func(context *gin.Context) {
		//支持内部和外部的重定向
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	engine.Run(":8000")
}
