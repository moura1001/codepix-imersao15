package grpc

import (
	"context"

	"github.com/moura1001/codepix/service/grpc/pb"
	"github.com/moura1001/codepix/service/usecase"
)

type PixKeyGrpcService struct {
	pixKeyUseCase usecase.PixKeyUseCase
	pb.UnimplementedPixServiceServer
}

func NewPixKeyGrpcService(useCase usecase.PixKeyUseCase) PixKeyGrpcService {
	return PixKeyGrpcService{
		pixKeyUseCase: useCase,
	}
}

func (p PixKeyGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	pixKey, err := p.pixKeyUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     pixKey.Id,
		Status: "created",
	}, nil
}

func (p PixKeyGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.pixKeyUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return nil, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.Id,
		Key:  pixKey.Key,
		Kind: string(pixKey.Kind),
		Account: &pb.Account{
			AccountId:     pixKey.AccountId,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankId,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}
