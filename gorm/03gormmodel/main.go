package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`
	Birthday     *time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`
	MemberNumber string `gorm:"unique;not null"` //唯一且非空
	Num          int    `gorm:"AUTO_INCREMENT"`  //自增
	Address      string `gorm:"index:add"`       //创建名为addr的索引
	IgnoreMe     int    `gorm:"-"`               //忽略本字段
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"` //手动
}

func (a Animal) TableName() string {
	return "qimi"
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "XXX_" + defaultTableName
	}
	//连接mysql数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed , error: ", err)
		return
	}
	defer db.Close()

	db.SingularTable(true) //禁用复数
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//使用user结构体来创建名叫useruseruser的表
	//db.Table("useruseruser").CreateTable(&User{})
}
