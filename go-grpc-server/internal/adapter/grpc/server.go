package grpc

import (
	"fmt"
	"github.com/jpmoraess/go-grpc-server/internal/port"
	"github.com/jpmoraess/grpc-proto/protogen/go/bank"
	"github.com/jpmoraess/grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	port         int
	server       *grpc.Server
	helloService port.HelloServicePort
	bankService  port.BankServicePort
	hello.HelloServiceServer
	bank.BankServiceServer
}

func NewGrpcAdapter(port int, helloService port.HelloServicePort, bankService port.BankServicePort) *GrpcAdapter {
	return &GrpcAdapter{
		port:         port,
		helloService: helloService,
		bankService:  bankService,
	}
}

func (a *GrpcAdapter) Run() {
	var err error
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen gRPC server on port %d : %v\n", a.port, err)
	}

	log.Printf("gRPC server listening on port %d\n", a.port)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)
	bank.RegisterBankServiceServer(grpcServer, a)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC on port %d : %v\n", a.port, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
