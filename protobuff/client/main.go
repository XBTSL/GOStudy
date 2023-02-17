package main

import (
	"context"
	"fmt"
	pb "github.com/my/repo/protobuff/sever/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//连接server端
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "xiongbin"})
	fmt.Println(resp.GetResponseMsg())
}
