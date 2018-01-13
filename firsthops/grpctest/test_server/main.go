package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/bsinou/go-proto3-training/firsthops/grpctest/test_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement test_service.CheckerServer.
type server struct{}

// GetServerInfo implements test_service.CheckerServer
func (s *server) GetServerInfo(ctx context.Context, in *pb.InfoRequest) (*pb.InfoList, error) {
	if in.UserName == "admin" {
		osName, err := os.Hostname()
		if err != nil {
			return nil, err
		} else {
			return &pb.InfoList{OsName: osName, LocalTime: time.Now().Format("yyyy-MM-dd HH:mm")}, nil
		}
	} else {
		return nil, fmt.Errorf("Unauthorised user...")
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServerCheckerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
