package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type student struct {
	sno, sname, ssex string
	sage             int
}

var db *sql.DB //是一个连接池对象
func main() {

	initDB()
	//queryOne("2016210867")
	//queryMany()
	//insert()
	//update()
	prepareSelect()
}
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tset"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("dsn %s is invalid , error : %v\n", dsn, err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("open %s failed , err : %v\n", dsn, err)
		return err
	}
	//设置数据库连接池最大连接数  如果超过了就会阻塞等待其他的程序
	db.SetMaxOpenConns(10)
	//最大空闲连接数
	db.SetMaxIdleConns(5)
	return nil
}
func queryOne(sno string) {
	var s student
	//查找
	row := db.QueryRow("select * from student where sno = ?", sno) //从连接池中拿一个连接出来去数据库查询单挑记录
	//扫描+释放连接  必须对rowobj调用scan方法
	row.Scan(&s.sno, &s.sname, &s.ssex, &s.sage)

	var sname string
	db.QueryRow("select sname from student where sno = ?", sno).Scan(&sname)
	fmt.Printf("student: %v \n", s)
	fmt.Println(sname)
}
func queryMany() {
	rows, err := db.Query("select sno , sname , ssex , sage from student ")
	if err != nil {
		fmt.Println("db query failed , error : %v\n", err)
		return
	}
	//非常重要 一定要关闭rows
	defer rows.Close()
	//循环取值
	for rows.Next() {
		var s student
		rows.Scan(&s.sno, &s.sname, &s.ssex, &s.sage)
		fmt.Println(s)
	}

}

func insert() {
	sql := "insert into student values ('2020170281','ljs','man',23)"
	exec, err := db.Exec(sql)
	if err != nil {
		fmt.Println("exec insert %s failed , err : %v\n", sql, err)
		return
	}
	//如果是插入数据的操作,能够拿到插入数据的id
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("get id failed ,err : %v\n", err)
		return
	}
	fmt.Println("id:", id)
}

func update() {
	exec, err := db.Exec("update student set sname = 'ljsnew' where sno='2020170281' ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())
}

//预处理方式select多条数据
func prepareSelect() {
	prepare, err := db.Prepare("select * from student where sno like ? ") //吧sql语句先发给mysql预处理一下
	if err != nil {
		return
	}

	rows, err := prepare.Query("%2016%") //后续只需要传值就行了
	if err != nil {
		return
	}

	defer prepare.Close()
	defer rows.Close()
	for rows.Next() {
		var s student
		rows.Scan(&s.sno, &s.sname, &s.ssex, &s.sage)
		fmt.Println(s)
	}
}
