package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	//初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	//加载页面
	engine.LoadHTMLGlob("./gin/12bookPractice/templates/*")
	//查询所有图书
	engine.GET("/book/list", bookListHandler)
	engine.GET("/book/new", func(context *gin.Context) {
		context.HTML(http.StatusOK, "new_book.html", gin.H{"code": 0})
	})
	engine.POST("/book/new", func(context *gin.Context) {
		title := context.PostForm("title")
		price := context.PostForm("price")
		atoi, err := strconv.Atoi(price)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"err": err})
			return
		}
		err = insertBook(title, int64(atoi))
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"err": err})
			return
		}
		context.Redirect(http.StatusPermanentRedirect, "/book/list")
		return
	})
	engine.GET("/book/delete", func(context *gin.Context) {
		//这是使用api传参 所以这边是接收不到值的
		context.Param("id")

		//实际上是使用的是url传参 所以要使用query来获得参数
		query := context.Query("id")

		atoi, err := strconv.Atoi(query)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"error": err})
			return
		}
		err1 := deleteBook(int64(atoi))
		if err1 != nil {
			context.JSON(http.StatusOK, gin.H{"error": err1})
			return
		}
		context.Redirect(http.StatusPermanentRedirect, "/book/list")
		return
	})
	engine.GET("/book/search", func(context *gin.Context) {
		query := context.Query("searchThing")
		book, err2 := queryBook(query)
		if err2 != nil {
			context.JSON(http.StatusOK, gin.H{"err": err2})
			return
		}
		context.HTML(http.StatusOK, "book_list.html", gin.H{"data": book})

	})
	engine.Run(":8000")
}

func bookListHandler(context *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 1, "msg": err})
		return
	}

	//返回数据
	context.HTML(http.StatusOK, "book_list.html", gin.H{"code": 0, "data": bookList})
}
