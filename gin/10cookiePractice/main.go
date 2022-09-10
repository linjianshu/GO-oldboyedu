package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWareAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("loginToken")
		if err != nil {
			//返回错误
			context.JSON(http.StatusOK, gin.H{"error": "StatusUnauthorized"})
			context.Abort()
			return
		}
		if cookie == "true" {
			context.Next()
		}
	}
}
func main() {
	engine := gin.Default()

	engine.GET("/login", func(context *gin.Context) {
		http.SetCookie(context.Writer, &http.Cookie{
			Name:  "loginToken",
			Value: "true",
		})
		context.String(http.StatusOK, "Login successful")
	})

	engine.GET("/home", MiddleWareAuth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	engine.Run(":8000")
}
