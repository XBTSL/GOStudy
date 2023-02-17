package main

import (
	"context"
	pb "github.com/my/repo/protobuff/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type personServer struct {
	pb.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *pb.PersonReq) (*pb.PersonRes, error) {
	return &pb.PersonRes{Age: req.Age, Name: req.Name}, nil
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
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
