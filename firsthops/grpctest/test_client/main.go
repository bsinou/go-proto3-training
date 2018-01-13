package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/bsinou/go-proto3-training/firsthops/grpctest/test_service"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerCheckerClient(conn)

	// Contact the server and print out its response.
	name := "admin"
	r, err := c.GetServerInfo(context.Background(), &pb.InfoRequest{UserName: name})
	if err != nil {
		log.Fatalf("could not getInfo: %v", err)
	}
	log.Printf("Greeting at %s from: %s", r.LocalTime, r.OsName)

	// Contact the server and print out its response.
	name = "user"
	r, err = c.GetServerInfo(context.Background(), &pb.InfoRequest{UserName: name})
	if err != nil {
		log.Fatalf("could not getInfo: %v", err)
	}
	log.Printf("Greeting at %s from: %s", r.LocalTime, r.OsName)
}
