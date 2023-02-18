package main

import (
	"context"
	"fmt"
	users "github.com/my/repo/grpcmysqlredis/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc 连接错误。。。。。")
		return
	}
	defer conn.Close()
	client := users.NewUserIdentifyClient(conn)
	resp, err := client.IdentifyUser(context.Background(), &users.UserEntry{Id: 2, Name: "赵悦", Password: "123456", Email: "2010216003@qq.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
