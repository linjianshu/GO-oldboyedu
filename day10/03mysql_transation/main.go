package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	initDB()
	transation()
}

func transation() {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}
	//执行多个sql操作
	_, err = tx.Exec("update student set sage = sage+5 where sno = '2016210867';")
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql1出错了,要回滚")
		return
	}
	_, err = tx.Exec("update xxx set sage = sage-5 where sno = '2020170281';")
	if err != nil {
		//如果出错了 要回滚
		tx.Rollback()
		fmt.Println("执行sql2出错了,要回滚")
		return
	}
	//如果执行成功了,
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错了,要回滚")
		return
	}
}

func initDB() (err error) {

	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tset")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
