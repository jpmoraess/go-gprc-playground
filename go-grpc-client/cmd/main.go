package main

import (
	"context"
	"fmt"
	"github.com/jpmoraess/go-grpc-client/internal/adapter/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalln("cannot connect to gRPC server: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	helloAdapter, err := hello.NewHelloAdapter(conn)
	if err != nil {
		log.Fatalln("cannot create HelloAdapter: ", err)
	}

	for i := range 10 {
		sayHello(helloAdapter, fmt.Sprintf("John %d", i))
	}
}

func sayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)
	if err != nil {
		log.Fatalln("cannot call SayHello: ", err)
	}
	log.Println("greeting: ", greet.Greet)
}
