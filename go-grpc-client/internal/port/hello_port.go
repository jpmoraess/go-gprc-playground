package port

import (
	"context"
	"github.com/jpmoraess/grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type HelloClientPort interface {
	SayHello(ctx context.Context, req *hello.HelloRequest, opts ...grpc.CallOption) (*hello.HelloResponse, error)
}
