package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	sqlInjectDemo("xxx'or 1=1#")
	sqlInjectDemo("xxx' union select * from student #")
}

var db *sqlx.DB

type student struct {
	SNO, SNAME, SSEX string
	SAGE             int
}

func sqlInjectDemo(name string) {
	//自己拼接sql语句
	sqlStr := fmt.Sprintf("select * from student where sname ='%s'", name)
	fmt.Println(sqlStr)

	var err error
	db, err = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/tset")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	var s []student
	err = db.Select(&s, sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
