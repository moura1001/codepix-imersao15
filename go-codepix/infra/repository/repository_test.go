package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura1001/codepix/domain/model"
	"github.com/moura1001/codepix/infra/db"
	"github.com/moura1001/codepix/service/factory"
	"github.com/stretchr/testify/require"
)

func Test_RepositoriesInsert(t *testing.T) {
	database, err := db.GetDBConnectionTest()
	require.Nil(t, err)

	pixKeyRepository := factory.NewPixKeyRepositoryDb(database)
	transactionRepository := factory.NewTransactionRepository(database)

	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	sourceAccountNumber := uuid.NewString()
	sourceOwnerName := "Moura1"
	sourceAccount, _ := model.NewAccount(sourceOwnerName, sourceAccountNumber, bank)

	destinationAccountNumber := uuid.NewString()
	destinationOwnerName := "Moura2"
	destinationAccount, _ := model.NewAccount(destinationOwnerName, destinationAccountNumber, bank)

	kind := "email"
	key := "email@email.com"
	pixKey, _ := model.NewPixKey(kind, key, sourceAccount)

	err = pixKeyRepository.AddBank(bank)
	require.Nil(t, err)
	bankDb, _, err := pixKeyRepository.FindBankByCode(bank.Code)
	require.Nil(t, err)
	require.Equal(t, bankDb.Id, bank.Id)
	require.Equal(t, bankDb.Code, bank.Code)
	require.Equal(t, bankDb.Name, bank.Name)

	err = pixKeyRepository.AddAccount(sourceAccount)
	require.Nil(t, err)
	sourceAccountDb, _, err := pixKeyRepository.FindAccountByNumber(bank.Code, sourceAccount.Number)
	require.Nil(t, err)
	require.Equal(t, sourceAccountDb.Id, sourceAccount.Id)
	require.Equal(t, sourceAccountDb.Number, sourceAccount.Number)
	require.Equal(t, sourceAccountDb.BankCode, bank.Code)

	err = pixKeyRepository.AddAccount(destinationAccount)
	require.Nil(t, err)
	destinationAccountDb, _, err := pixKeyRepository.FindAccountByNumber(bank.Code, destinationAccount.Number)
	require.Nil(t, err)
	require.Equal(t, destinationAccountDb.Id, destinationAccount.Id)
	require.Equal(t, destinationAccountDb.Number, destinationAccount.Number)
	require.Equal(t, destinationAccountDb.BankCode, bank.Code)

	pixKeyDb, err := pixKeyRepository.RegisterKey(pixKey)
	require.Nil(t, err)
	require.Equal(t, pixKeyDb.Id, pixKey.Id)
	require.Equal(t, pixKeyDb.Kind, pixKey.Kind)
	require.Equal(t, pixKeyDb.Status, model.PIX_KEY_STATUS_ACTIVE)
	require.Equal(t, pixKeyDb.Status, pixKey.Status)

	amount := 7.77
	descriptionTransaction := "My description"
	transaction1, _ := model.NewTransaction(destinationAccount, amount, pixKey, descriptionTransaction)
	err = transactionRepository.Register(transaction1)
	require.Nil(t, err)
	transaction1.Confirm()
	err = transactionRepository.Save(transaction1)
	require.Nil(t, err)
	transaction1Db, err := transactionRepository.FindById(transaction1.Id)
	require.Nil(t, err)
	require.Equal(t, transaction1Db.Amount, transaction1.Amount)
	require.Equal(t, transaction1Db.Status, model.TRANSACTION_STATUS_CONFIRMED)

	cancelDescription := "Error"
	transaction1.Cancel(cancelDescription)
	err = transactionRepository.Save(transaction1)
	require.Nil(t, err)
	transaction1Db, err = transactionRepository.FindById(transaction1.Id)
	require.Nil(t, err)
	require.Equal(t, transaction1Db.Status, model.TRANSACTION_STATUS_CANCELLED)
	require.Equal(t, transaction1Db.CancelDescription, cancelDescription)
}
