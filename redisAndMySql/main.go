package main

import (
	"fmt"
	"time"

	//此处注意“_”表示引用mysql函数中init的方法而无需使用函数
	_ "github.com/go-sql-driver/mysql"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	//对应id表字段
	Id int `db:"id"`
	//对应name表字段
	Name string `db:"username"`
	//对应age表字段
	Password string `db:"password"`
	//对应rmb表字段
	Email string `db:"email"`
}

func main() {
	var cmd string

	for {
		fmt.Println("请输入命令:")
		fmt.Scan(&cmd)
		fmt.Println("你输入的是:", cmd)

		switch cmd {
		case "getall":
			GetAll()
		default:
			fmt.Println("不能识别的命令")
		}

		fmt.Println()
	}

}

func GetAll() {
	//先看看redis里有没有数据
	setDb := redis.DialDatabase(0)
	setPasswd := redis.DialPassword("tomjack")              //密码
	timeout := redis.DialConnectTimeout(5 * time.Second)    //连接超时时间
	readTimeout := redis.DialReadTimeout(5 * time.Second)   //读超时时间
	writeTimeout := redis.DialWriteTimeout(5 * time.Second) //写超时时间get

	conn, _ := redis.Dial("tcp", "localhost:8888", setDb, setPasswd, timeout, readTimeout, writeTimeout)
	defer conn.Close()
	reply, err := conn.Do("lrange", "mlist", 0, 10)
	pkeys, _ := redis.Strings(reply, err)
	fmt.Println(pkeys)

	if len(pkeys) > 0 {
		//如果有
		fmt.Println("从redis获得数据")

		// 从redis里直接读取
		for _, key := range pkeys {
			retStrs, _ := redis.Strings(conn.Do("hgetall", key))
			fmt.Println(retStrs)
			fmt.Printf("{%s %s %s}\n", retStrs[1], retStrs[3], retStrs[5])
		}

	} else {
		//如果没有
		fmt.Println("从mysql获得数据")

		//查询数据库
		db, err := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/test")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		var persons []Person
		db.Select(&persons, "select * from users")
		fmt.Println(persons)

		//写入redis并且设置过期时间
		for _, p := range persons {
			//将p以hash形式写入redis
			_, e1 := conn.Do("hmset", p.Id, "name", p.Name, "age", p.Password, "rmb", p.Email)

			//将这个hash的key加入mlist
			_, e2 := conn.Do("rpush", "mlist", p.Id)

			//设置过期时间
			_, e3 := conn.Do("expire", p.Id, 60)
			_, e4 := conn.Do("expire", "mlist", 60)

			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				fmt.Println(p.Name, "写入失败", e1, e2, e3, e4)
			} else {
				fmt.Println(p.Name, "写入成功")
			}
		}
	}
}
