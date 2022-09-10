package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"time"
)

type Student struct {
	Sno       string
	Sname     string
	Ssex      string
	Sbirthday time.Time
	Class     string
	Reserve1  string
	Reserve2  string
	Image     []byte
}

type Score struct {
	Sno     string
	Cno     string
	Degree1 string
}

func main() {
	//db, err := gorm.Open("mssql", "sqlserver://sa:123456@localhost:1433?database=Test")
	db, err := gorm.Open("mssql", "sqlserver://sa:123456@localhost?database=Test&connection+timeout=30")
	if err != nil {
		fmt.Println("open mssql failed . error : ", err)
		return
	}
	defer db.Close()

	//db.AutoMigrate(&Student{})
	var student Student
	db.First(&student)
	fmt.Println(student)

	var score Score
	first := db.First(&score)
	fmt.Println(first.RowsAffected)
	fmt.Println(score)
}
