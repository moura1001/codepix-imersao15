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

func (p PixKeyGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.RegisterCreatedResult, error) {
	pixKey, err := p.pixKeyUseCase.RegisterKey(in.Key, in.Kind, in.AccountNumber, in.BankCode)
	if err != nil {
		return &pb.RegisterCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterCreatedResult{
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
			AccountNumber: pixKey.Account.Number,
			BankCode:      pixKey.Account.BankCode,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func (p PixKeyGrpcService) RegisterAccount(ctx context.Context, in *pb.Account) (*pb.RegisterCreatedResult, error) {
	err := p.pixKeyUseCase.RegisterAccount(in.OwnerName, in.AccountNumber, in.BankCode)
	if err != nil {
		return &pb.RegisterCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterCreatedResult{
		Status: "created",
	}, nil

}

func (p PixKeyGrpcService) RegisterBank(ctx context.Context, in *pb.Bank) (*pb.RegisterCreatedResult, error) {
	err := p.pixKeyUseCase.RegisterBank(in.BankCode, in.BankName)
	if err != nil {
		return &pb.RegisterCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterCreatedResult{
		Status: "created",
	}, nil

}
