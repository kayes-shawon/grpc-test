package main

import (
	"context"
	"github.com/kayes-shawon/grpc-test/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error: %v",err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	response, err := c.SayHello(context.TODO(), &proto.HelloRequest{Name:"Kayes"})
	if err != nil {
		log.Fatalf("Error: %v",err)
	}
	log.Printf("Response from server: %s", response.Message)
}