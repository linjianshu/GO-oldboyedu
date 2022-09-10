package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type student struct {
	SNO, SNAME, SSEX string
	SAGE             int
}

func main() {
	initDB()
	var s1 student
	var s = make([]student, 0, 10)
	err := db.Get(&s1, "select * from student;")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	err = db.Select(&s, "select * from student")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)

}
func initDB() (err error) {

	db, err = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/tset")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
