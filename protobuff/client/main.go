package main

import (
	"context"
	"fmt"
	pb "github.com/my/repo/protobuff/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//连接server端
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	//建立连接
	client := pb.NewSearchServiceClient(conn)
	//执行rpc调用
	resp, _ := client.Search(context.Background(), &pb.PersonReq{Age: 11, Name: "tom"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
