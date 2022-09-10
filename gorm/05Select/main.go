package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model // id createat update at deleteat
	Name       string
	Age        int64
}

func (u User) TableName() string {
	return "SelectUser"
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/tset?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql open failed ,error : ", err)
		return
	}
	defer db.Close()

	//自动迁移 或者对应
	db.AutoMigrate(&User{})

	////创建
	//db.Create(&User{Age: 18,Name: "ljs"})
	//
	//user := User{Age: 28, Name: "jwt"}
	//db.Create(&user)

	//查询 根据主键查询第一条
	user := User{}
	db.First(&user)
	//select * from users order by id limit 1
	fmt.Println(user)

	//随机获取一条
	u := new(User)
	db.Take(u)
	fmt.Println(*u)

	//根据主键查询最后一条记录
	u2 := new(User)
	db.Last(u2)
	fmt.Println(*u2)

	//查询所有的记录
	var users []User
	db.Debug().Find(&users)
	fmt.Println(users)

	//查询特定某条记录 (仅当主键为整型的时候可用)
	u3 := new(User)
	db.Find(&u3, 2)

	//普通sql查询
	//get first
	db.Where("name=?", "ljs").First(&user)
	fmt.Println(user)

	db.Where("name like ?", "%%%").Find(&users)
	fmt.Println(users)

	db.Where("updated_at < ?", time.Now()).Find(&users)
	fmt.Println(users)

	//struct && map 查询
	db.Where(&User{Name: "ljs", Age: 18}).First(&user)
	fmt.Println(user)

	//map
	db.Where(map[string]interface{}{"name": "ljs", "age": 18}).Find(&user)
	fmt.Println(user)

	//主键的切片
	db.Where([]int{1, 2}).Find(&users)
	fmt.Println(users)

	//not查询
	db.Not("name", "jinzhu").First(&user)
	fmt.Println(user, "notnot")

	//not in
	db.Not("name", []string{"linjianshu", "jwt"}).Find(&user)
	fmt.Println(user)

	//not in primary
	u4 := new(User)
	db.Debug().Not([]int64{1}).Find(u4)
	fmt.Println(u4, "not in primary")

	//plain sql
	// = 和 ? 之间一定要有空格
	u5 := new(User)
	db.Debug().Not("name = ?", "ljs").First(u5)
	fmt.Println(*u5)

	//struct
	u6 := new(User)
	u6.Name = "ljs"
	u6.Age = 18
	u7 := new(User)
	db.Debug().Not(&u6).First(&u7)
	fmt.Println(u7, " not in struct ")

	//or条件
	i := new([]User)
	db.Where("name = ?", "ljs").Or("name = ?", "jwt").Find(i)
	fmt.Println(i, " or condition ")

	//struct or
	i2 := new([]User)
	db.Where(User{Name: "ljs"}).Or(User{Age: 28}).Find(i2)
	fmt.Println(i2, "or struct")

	//map
	i3 := new([]User)
	db.Where(map[string]interface{}{"name": "ljs"}).Or(map[string]interface{}{"age": 18}).Find(i3)
	fmt.Println(i3, " or map ")

	//内联条件
	u8 := new(User)
	db.Debug().First(u8, 2)
	fmt.Println(u8)

	//plan sql
	u9 := new(User)
	db.Find(u9, "name = ?", "ljs")
	fmt.Println(u9)

	//struct map ...

	//firstOrInit
	//未找到
	u10 := new(User)
	db.FirstOrInit(u10, User{Name: "fuck no existing"})
	fmt.Println(u10)

	//attr
	db.Attrs(User{Age: 100}).FirstOrInit(u10, User{Name: "fuck no existing"})
	fmt.Println(u10)

	//assign
	u11 := new(User)
	db.Assign(User{Age: 1000}).Where(User{Name: "jwt"}).FirstOrInit(u11)
	fmt.Println(u11)

	//order
	i4 := new([]User)
	db.Order("age desc , name").Find(i4)
	fmt.Println(i4)

	i5 := new([]User)
	db.Order("age desc").Order("name").Find(i5)
	fmt.Println(i5)

	//覆盖了
	i6 := new([]User)
	db.Order("age desc").Find(i6).Order("age", true).Find(i6)
	fmt.Println(i6)

	//数量
	u12 := new(User)
	db.Limit(1).Find(u12)
	//-1 取消limit条件
	//i7 := new([]User)
	//db.Limit(2).Find(i7).Limit(-1).Find()

	//offset 偏移 也就是takeWhile
	u13 := new(User)
	db.Offset(1).First(u13)
	fmt.Println(u13)

	//count
	i7 := new(int)
	println(db.Table("selectuser").Count(i7))
	fmt.Println(*i7)

}
