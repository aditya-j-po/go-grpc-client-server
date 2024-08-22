package main

import (
	"log"
	"net"

	pb "github.com/aditya-j-po/go-grpc-client-server/proto"
	"google.golang.org/grpc"
)

// define the port
const (
	port = ":8080"
)

// this is the struct to be created, pb is imported upstairs
type helloServer struct {
	pb.GreetServiceServer
}

func SetupGrpcServer() error {
	// listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	return err
}

func main() {

	err := SetupGrpcServer()
	if err != nil {
		log.Fatalf("Main had error: %v", err)
	}
}
