package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//go连接mysql实例
	//数据库源信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/tset"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {
		log.Printf("dsn %s invalid , err : %v\n", dsn, db)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("open %s failed , err : %v\n", dsn, err)
		return
	}
	fmt.Println(db, "连接数据库成功!")

}
