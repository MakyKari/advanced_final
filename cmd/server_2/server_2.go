package main

import (
	"context"
	"log"
	"net"

	"github.com/dungtc/grpc-playground/simple/helloworld"
	"google.golang.org/grpc"
)

var (
	address = "10001"
)

type handler struct {
}

func (h *handler) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	passwd := "123456"
	if req.GetName() != passwd {
		return &helloworld.HelloReply{
			Message: "Invalid password",
		}, nil
	}
	return &helloworld.HelloReply{
		Message: "Success",
	}, nil
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
