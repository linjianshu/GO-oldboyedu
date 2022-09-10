package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("TryCookie", func(context *gin.Context) {
		//获取客户端是否携带cookie
		//cookies := context.Request.Cookies()
		//一样的
		cookie, err := context.Cookie("key_cookie")
		cookies := context.Request.Cookies()
		if err != nil {
			fmt.Println("notset")
			//设置cookie
			//maxAge int 单位为秒
			//path cookie所在目录
			//domain  string 域名
			//secure 是否只能通过https访问
			//httponly bool  是否允许别人通过js获取自己的cookie
			//context.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
			//不知道为什么上面的方法行不通  好像是参数设置的有问题
			http.SetCookie(context.Writer, &http.Cookie{
				Name:     "key_cookie",
				Value:    "value_cookie",
				Domain:   "localhost",
				MaxAge:   60,
				Secure:   false,
				HttpOnly: true,
			})

			return
		}

		fmt.Println("cookie的值是: ", cookie)
		fmt.Println(cookies)
	})
	engine.Run(":8000")

}
