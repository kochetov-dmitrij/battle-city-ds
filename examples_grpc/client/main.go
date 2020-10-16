package main

import (
	"context"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewComsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Connected a client to %s", address)

	// Add a new product and get its ID
	newProduct := &pb.Message{
		TankPosition:    nil,
		BulletPosition:  nil,
		BulletDirection: pb.Message_UP,
		Action:          nil,
	}
	_, err = c.AddMessage(ctx, newProduct)
	if err != nil {
		log.Fatalf("Could not send the product: %v", err)
	}
}
