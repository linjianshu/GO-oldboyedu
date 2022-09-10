package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Sno   string
	Sname string
	Ssex  string
	Sage  int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed , error :", err)
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移 (把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{Sno: "1", Sname: "jwt", Ssex: "nan", Sage: 18}
	result := db.Create(&u1)
	fmt.Println(result.RowsAffected)

	//查询
	var student UserInfo
	db.First(&student)
	fmt.Println(student)

	//更新
	db.Model(&student).Update("Ssex", "nv")
	//此时已经修改了结构体的属性了
	fmt.Println(student)

	//删除
	db.Delete(&student)
	//这样不行 虽然表里删除了 但是没有清空student这个实体的值
	db.First(&student)
	var studentD UserInfo
	db.First(&studentD)
	fmt.Println(studentD)

}
