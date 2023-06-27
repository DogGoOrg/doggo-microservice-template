package main

import (
	"context"
	"doggo-microservice-template/internal/db"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	// PACKAGE.UnimplementedAccountServer add one
	Database *gorm.DB
}

// ping method implementation
func (receiver *server) Ping(ctx context.Context, in *PACKAGE.PingRequest) (*PACKAGE.PingResponse, error) {
	fmt.Println("OK")
	return &PACKAGE.PingResponse{
		Status: "OK",
	}, nil
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := new(server)

	db := db.InitDB()

	serv.Database = db

	db.AutoMigrate()

	s := grpc.NewServer()

	PACKAGE.RegisterAccountServer(s, serv)
	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
