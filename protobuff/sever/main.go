package main

import (
	"context"
	helloservice "github.com/my/repo/protobuff/sever/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type personServer struct {
	helloservice.UnimplementedSayHelloServer
}

func (*personServer) Search(ctx context.Context, req *pb.PersonReq) (*pb.PersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*personServer) SearchIn(pb.SearchService_SearchInServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchIn not implemented")
}
func (*personServer) SearchOut(*pb.PersonReq, pb.SearchService_SearchOutServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOut not implemented")
}
func (*personServer) SearchIO(pb.SearchService_SearchIOServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchIO not implemented")
}

func main() {
	//开启端口
	listener, _ := net.Listen("tcp", ":8080")
	//创建grpc服务
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &personServer{})
}
