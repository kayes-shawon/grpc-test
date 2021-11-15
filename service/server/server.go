package main

import (
	"fmt"
	"github.com/kayes-shawon/grpc-test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)


func main() {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8080)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listed %v", err)
	}

	s := new(proto.Server)

	grpcServer := grpc.NewServer()

	proto.RegisterGreeterServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("FAILED TO SERVE: %s", err)
	}

}
