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
	return "DeleteUser"
}
func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql open failed ,error : ", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//db.Debug().Create(&User{
	//	Name:   "jwt",
	//	Age:    20,
	//	Active: false,
	//})
	//db.Debug().Create(&User{
	//	Name:   "lje",
	//	Age:    18,
	//	Active: false,
	//})

	//删除
	//没有主键的时候很可怕 会删除所有的
	//user :=User{}
	////user.ID  = 1
	//user.Name ="qimi"
	//db.Debug().Delete(&user)

	////选择性批量删除
	//i := new([]User)
	////使用结构体的时候 传入的参数如果是零值的话 不会作为筛选条件
	//db.Debug().Where(User{Active: false}).Find(i)
	//fmt.Println(i)
	////delete i 不会将i作为批量删除的筛选条件
	////db.Debug().Delete(i)
	//
	//db.Debug().Where(map[string]interface{}{"active":false}).Find(i)
	//fmt.Println(i)
	//
	////这样删除才可以 有条件删除
	//db.Debug().Where(map[string]interface{}{"active":false}).Find(i).Delete(User{})
	//
	//i2 := new(int)
	//fmt.Println(db.Debug().Model(User{}).Count(i2))

	//sql删除
	//db.Debug().Delete(User{},"active like 1")

	//就想查到被软删除的内容
	//i := new([]User)
	//db.Unscoped().Debug().Where("deleted_at is not null").Find(i)
	//db.Unscoped().Debug().Where("deleted_at is not null").Find(i).UpdateColumn("deleted_at",nil)

	//db.Debug().Where("deleted_at is not null").Find(i)

	//物理删除 硬删除
	db.Debug().Unscoped().Where("name = 'lje'").Delete(User{})

}
