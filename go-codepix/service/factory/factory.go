package factory

import (
	"github.com/moura1001/codepix/infra/repository"
	"github.com/moura1001/codepix/service/usecase"
	"gorm.io/gorm"
)

func NewPixKeyRepositoryDb(database *gorm.DB) repository.PixKeyRepositoryDb {
	return repository.PixKeyRepositoryDb{Db: database}
}

func NewPixKeyUseCase(database *gorm.DB) usecase.PixKeyUseCase {
	pixKeyRepository := NewPixKeyRepositoryDb(database)
	pixKeyUseCase := usecase.NewPixKeyUseCase(pixKeyRepository)
	return pixKeyUseCase
}

func NewTransactionRepository(database *gorm.DB) repository.TransactionRepositoryDb {
	return repository.TransactionRepositoryDb{Db: database}
}

func NewTransactionUseCase(database *gorm.DB) usecase.TransactionUseCase {
	pixKeyRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	return usecase.NewTransactionUseCase(transactionRepository, pixKeyRepository)
}
