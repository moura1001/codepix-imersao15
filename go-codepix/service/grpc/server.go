package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/moura1001/codepix/infra/repository"
	"github.com/moura1001/codepix/service/grpc/pb"
	"github.com/moura1001/codepix/service/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// pix key grpc sevice
	pixKeyRepository := repository.PixKeyRepositoryDb{Db: database}
	pixKeyUseCase := usecase.NewPixKeyUseCase(pixKeyRepository)
	pixKeyGrpcService := NewPixKeyGrpcService(pixKeyUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixKeyGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot listen grpc server on port %d. Details: '%s'", port, err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start grpc server. Details: '%s'", err)
	}
}
