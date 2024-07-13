package main

import (
	"github.com/jpmoraess/go-grpc-server/internal/adapter/grpc"
	"github.com/jpmoraess/go-grpc-server/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &application.HelloService{}

	grpcAdapter := grpc.NewGrpcAdapter(50051, hs)

	grpcAdapter.Run()
}
