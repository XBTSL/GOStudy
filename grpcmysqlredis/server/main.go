package main

import (
	"context"
	sql2 "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	users "github.com/my/repo/grpcmysqlredis/server/pb"
	"google.golang.org/grpc"
	"net"
)

var Db *sql2.DB

type user struct {
	id       int32
	name     string
	password string
	email    string
}

func initMySql(dsn string) (err error) {
	Db, err = sql2.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	Db.SetMaxOpenConns(10)
	Db.SetConnMaxIdleTime(10)
	fmt.Println("mysql 连接成功。。。。")
	return nil
}

func selectData(id int32) (result user, err error) {
	sqlStr := "select * from users where id = ?"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		err = rows.Scan(&result.id, &result.name, &result.password, &result.email)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

type userServer struct {
	users.UnimplementedUserIdentifyServer
}

func (tem *userServer) IdentifyUser(ctx context.Context, entry *users.UserEntry) (*users.UserResp, error) {
	tempid := entry.Id
	result, err := selectData(tempid)
	if err != nil {
		fmt.Println(err)
		return &users.UserResp{}, err
	}
	if result.password == entry.Password {
		return &users.UserResp{Name: entry.Name, Words: "welcomee to world......."}, nil
	}
	return &users.UserResp{}, nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	err := initMySql(dsn)
	//defer Db.Close()
	if err != nil {
		fmt.Println("mysql 连接失败。。。。")
		return
	}
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	users.RegisterUserIdentifyServer(server, &userServer{})
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
