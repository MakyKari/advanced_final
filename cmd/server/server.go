package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/dungtc/grpc-playground/simple/helloworld"
	"google.golang.org/grpc"
)

var (
	address = "10000"
)

type handler struct {
}

func (h *handler) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	message :=  req.GetName()

	addr := "localhost:10001"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: message})
	if err != nil {
		log.Fatal(err)
	}

	if res.GetMessage() == "Success" {
		return &helloworld.HelloReply{
			Message: "Welcome to the server!",
		}, nil
	} else {
		return &helloworld.HelloReply{
			Message: "Invalid password",
		}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &handler{})

	log.Printf("GRPC Server Listening in %s", address)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
