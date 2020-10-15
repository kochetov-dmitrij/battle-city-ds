package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	//products = map[uint32]*pb.Product{}
	lis, _ := net.Listen("tcp", port)
	server := grpc.NewServer()
	log.Printf("Launched a server on port %s", port)
	pb.RegisterComsService(server,
		&pb.ComsService{AddMessage: AddMessage})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func AddMessage(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	log.Printf("All products on the server: %v", msg.BulletDirection)
	return &empty.Empty{}, nil
}
