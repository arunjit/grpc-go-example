package main

import (
	"flag"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/arunjit/grpc-go-example/proto"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("New request: %v", in.String())
	return &pb.HelloResponse{Greeting: "Hello, " + in.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("boo")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
