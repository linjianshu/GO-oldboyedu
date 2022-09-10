package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Employee 1.定义模型 使用指针
type Employee struct {
	ID   int64
	Name *string `gorm:"default:'小王子'"`
	Age  int64
}

// Employee1 1.使用scanner/valuer
type Employee1 struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

// TableName 定义已存在表和结构体的映射关系
func (e Employee1) TableName() string {
	return "employees"
}

func main() {

	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed , error: ", err)
		return
	}

	defer db.Close()

	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&Employee{})

	//a:="ljs"
	//3.创建
	//employee := Employee{
	//	Age:  38,
	//	Name: 	new(string),
	//}
	//
	//employee1 := Employee{
	//	Age:  48,
	//	Name: 	&a,
	//}

	db.CreateTable()
	employee2 := Employee1{
		Age: 58,
		Name: sql.NullString{
			String: "",
			Valid:  true,
		},
	}
	//db.NewRecord(&employee) //判断主键是否为空
	//db.Debug().Create(&employee)
	//db.Debug().Create(&employee1)
	db.Debug().Create(&employee2)
	//fmt.Println(db.NewRecord(&employee))

}
