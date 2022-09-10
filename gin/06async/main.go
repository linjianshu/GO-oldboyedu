package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	engine := gin.Default()
	//1.异步
	engine.GET("/long_async", func(context *gin.Context) {
		//需要搞一个只读的副本
		copyContext := context.Copy()
		//模仿异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行..." + copyContext.Request.URL.Path)
		}()
	})

	//2.同步
	engine.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行..." + context.Request.URL.Path)
	})
	engine.Run(":8000")
}
