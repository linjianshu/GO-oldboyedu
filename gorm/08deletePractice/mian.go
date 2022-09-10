package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Beauty 首字母大写 否则娶不到值
type Beauty struct {
	Id           int       `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Sex          string    `gorm:"column:sex"`
	Borndate     time.Time `gorm:"column:borndate"`
	Phone        string    `gorm:"column:phone"`
	Photo        []byte    `gorm:"column:photo"`
	Boyfriend_id int       `gorm:"column:boyfriend_id"`
}

// TableName 防止数据库中本来就有 却重新生成一张表
func (b Beauty) TableName() string {
	return "beauty"
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/girls?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("girls open failed , error : ", err)
		return
	}
	defer db.Close()

	//db.AutoMigrate(&Beauty{})

	i := new([]Beauty)
	db.Find(i)
	fmt.Println(i)

	b := new(Beauty)
	db.Where("id = ? ", 10).First(b)
	fmt.Println(b)

	b3 := new(Beauty)
	db.Where(Beauty{
		Id: 10,
	}).Find(&b3)
	fmt.Println(b3)

	b3.Phone = "13063336291"
	db.Debug().Save(&b3)

	b2 := new(Beauty)
	db.Where("id = 10").First(b2)
	fmt.Println(b2)

	b2.Phone = "110"
	db.Debug().Model(&b2).Update("phone", "110")

	db.Debug().Delete(&b2)

}
