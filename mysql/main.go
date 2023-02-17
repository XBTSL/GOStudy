package main

import (
	sql2 "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	userName string
	password string
	email    string
}

var (
	Db *sql2.DB
)

func initDB(dsn string) (err error) {
	Db, err = sql2.Open("mysql", dsn)
	// open函数只是验证格式是否正确,不会校验账号密码是否正确,并不是创建数据库连接
	if err != nil {
		return
	}
	// 尝试与数据库建立连接，即校验dsn是否正确
	err = Db.Ping()
	if err != nil {
		return
	}
	//SetMaxOpenConns设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
	//SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。
	Db.SetMaxOpenConns(10)
	Db.SetConnMaxIdleTime(10)
	fmt.Println("数据库链接成功\n")
	return
}

func addUser(id int, username, password, email string) {
	str := "insert into users(id,username,password,email) values(?,?,?,?)"
	ret, err := Db.Exec(str, id, username, password, email)
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	theUser, err := ret.LastInsertId() //新插入的数据
	if err != nil {
		fmt.Println("get lastInsert fail:", err)
		return
	}
	fmt.Println("insert success,the id is ", theUser)
}

func selectManyData(id int) {
	sqlStr := "select * from users where id > ?"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Println("数据查询失败:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tem user
		err := rows.Scan(&tem.id, &tem.userName, &tem.password, &tem.email)
		if err != nil {
			fmt.Println("数据读取失败. err:", err)
			return
		}
		// 输出数据
		fmt.Printf("id:%d name:%s passqord:%s email:%s\n", tem.id, tem.userName, tem.password, tem.email)
	}
}

func updateData(id int, password string) {
	salStr := "update users set password=? where id=?"
	ret, err := Db.Exec(salStr, password, id)
	if err != nil {
		fmt.Println("数据更新失败.err:", err)
	}
	rowNum, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取影响的行数失败. err:", err)
		return
	}
	// 打印影响的行数
	fmt.Printf("修改成功，影响行数:%d\n", rowNum)
}

func deleteData(id int) {
	sqlStr := "delete from users where id = ?"
	ret, err := Db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("数据删除失败, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("获取受影响行数失败, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功，影响行数:%d\n", n)

}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	err := initDB(dsn)
	defer Db.Close()
	if err != nil {
		fmt.Println("数据库初始化失败:", err)
		return
	}
	addUser(0001, "熊彬", "123456", "2010216003@qq.com")
	//addUser(0002, "赵悦", "123456", "2010216003@qq.com")
	selectManyData(0)
	updateData(1, "999999")
	selectManyData(0)
	deleteData(1)
	selectManyData(0)
}
