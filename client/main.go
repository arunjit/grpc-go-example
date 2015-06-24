package main

import (
	"flag"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/arunjit/grpc-go-example/proto"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "Server address host:port")
)

func main() {
	conn, err := grpc.Dial(*addrFlag)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Arunjit"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Greeting)
}
