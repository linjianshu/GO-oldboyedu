package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

//多种响应方式
func main() {
	//1.创建路由
	//2.默认使用了2个中间件logger recovery
	engine := gin.Default()

	//1.json
	engine.GET("/someJson", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "someJson", "Status": 200})
	})

	//2.结构体响应
	engine.GET("someStruct", func(context *gin.Context) {
		context.JSON(http.StatusOK, struct {
			Name, Message string
			Number        int
		}{"root", "message", 123})
	})

	//3.XML响应
	engine.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "xml"})
	})

	//4.yaml响应
	engine.GET("/someYaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"name": "YAML"})
	})

	//5.protobuf格式 副歌开发的高校存储读取的工具
	engine.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{1, 2}
		//定义返回数据
		label := "protobuf"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		context.ProtoBuf(http.StatusOK, data)
	})

	engine.Run(":8000")
}
