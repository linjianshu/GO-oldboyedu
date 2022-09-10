package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func middleTimer() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		context.Next()
		fmt.Println("用时: ", time.Now().Sub(now))
	}
}
func main() {
	engine := gin.Default()
	engine.Use(middleTimer())
	group := engine.Group("/timer")
	group.GET("/countTimer", func(context *gin.Context) {
		time.Sleep(3 * time.Second)
	})
	group.GET("/counTimer1", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
	})
	engine.Run(":8000")
}
