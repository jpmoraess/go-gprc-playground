package hello

import (
	"context"
	"github.com/jpmoraess/go-grpc-client/internal/port"
	"github.com/jpmoraess/grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := hello.NewHelloServiceClient(conn)
	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(ctx context.Context, name string) (*hello.HelloResponse, error) {
	request := &hello.HelloRequest{
		Name: name,
	}

	greet, err := a.helloClient.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("error on SayHello : %v", err)
	}

	return greet, nil
}
