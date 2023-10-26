package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/moura1001/codepix/service/factory"
	"github.com/moura1001/codepix/service/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// pix key grpc sevice
	pixKeyUseCase := factory.NewPixKeyUseCase(database)
	pixKeyGrpcService := NewPixKeyGrpcService(pixKeyUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixKeyGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot listen grpc server on port %d. Details: '%s'", port, err)
	}

	log.Printf("gRPC server has been started on port %d\n", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start grpc server. Details: '%s'", err)
	}
}
