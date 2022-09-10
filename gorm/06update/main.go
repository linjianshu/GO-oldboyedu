package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name   string
	Age    int
	Active bool
}

func (u User) TableName() string {
	return "UpdateUser"
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql open failed ,error : ", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//db.Create(&User{Name: "qimi", Age: 18, Active: true})
	//db.Create(&User{Name: "jinzhu", Age: 20, Active: false})

	//var user = User{}
	//db.First(&user)
	//fmt.Println(user)
	//
	////更新
	//user.Name ="林建树"
	//user.Age=99
	////save 默认修改所有字段
	//db.Debug().Save(&user)
	//
	//db.Debug().Model(&user).Update("name","xwz")
	//
	//i := new([]User)
	//db.Debug().Where("name = ?","xwz").Or("active = ?",false).Find(i)
	//fmt.Println(i)
	//
	////更新单个属性
	//db.Model(&i).Where("age = ?",99).Update("name","hello")
	////更新完后还是不会变化的 原来查到的 i
	//fmt.Println(i)
	//
	////使用map更新多个属性 , 只会更新其中有变化的属性
	//i2 := new([]User)
	//db.Find(i2)
	//fmt.Println(i2)
	//
	//db.Debug().Model(i2).Where(User{Age: 99}).Updates(map[string]interface{}{"age":100,"active":false})
	//
	////使用struct更新多个属性,只会更新其中有变化且非零值的字段
	//i3 := new([]User)
	//db.Debug().Model(i3).Updates(User{Active: true})
	//
	////警告：当使用struct更新时，gorm只会更新那些非零值的字段
	////对于下面的操作，不会发生任何更新，因为false都是那些非零值的指端
	//db.Debug().Model(i3).Updates(User{Active: false})
	//
	////更新选定字段 select
	//i4 := new(User)
	////只适合于传入的参数是user 而不是 []user
	//db.Debug().Model(i4).Select("active").Updates(map[string]interface{}{"name":"hello1","active":false})
	//
	////omit 忽略
	//i5 := new(User)
	////只适合于传入的参数是user 而不是 []user
	//db.Debug().Model(i5).Omit("name").Updates(map[string]interface{}{"name":"hello","active":true})
	//
	////无hook更新  不会更新update_at
	//u := new(User)
	////db.Debug().Where("age= ?",100).First(u)
	////db.Debug().Where(map[string]interface{}{"age":100}).First(u)
	//db.Debug().Where(User{Age: 100}).First(u)
	//fmt.Println(u)
	//db.Debug().Model(u).UpdateColumn("name","ljs")
	//
	//i6 := new([]User)
	//db.Debug().Where(map[string]interface{}{"active":true}).Find(i6)
	////model在这里的作用并不能限制固定的数量 只能起到映射的作用
	//println(db.Debug().Model(i6).UpdateColumns(User{Active: true, Age: 18}).RowsAffected)

	//批量更新

	//使用sql表达式更新
	i7 := new([]User)
	db.Debug().Model(User{}).Where("active = ?", true).Find(&i7)
	//model不能作为限制输入参数的条件
	db.Debug().Model(i7).Update("age", gorm.Expr("age * ? + ?", 1, 1))

	db.Debug().Model(User{}).Where("active = ?", true).Update("age", gorm.Expr("age * ? + ?", 1, 1))
}
