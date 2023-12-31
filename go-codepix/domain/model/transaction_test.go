package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura1001/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTransaction(t *testing.T) {
	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	sourceAccountNumber := uuid.NewString()
	sourceOwnerName := "Moura"
	sourceAccount, _ := model.NewAccount(sourceOwnerName, sourceAccountNumber, bank)

	destinationAccountNumber := uuid.NewString()
	destinationOwnerName := "Moura"
	destinationAccount, _ := model.NewAccount(destinationOwnerName, destinationAccountNumber, bank)

	kind := "email"
	key := "email@email.com"
	pixKey, _ := model.NewPixKey(kind, key, sourceAccount)

	require.NotEqual(t, sourceAccount.Id, destinationAccount.Id)

	amount := 7.77
	statusTransaction := model.TRANSACTION_STATUS_PENDING
	descriptionTransaction := "My description"
	transaction, err := model.NewTransaction(destinationAccount, amount, pixKey, descriptionTransaction)

	require.Nil(t, err)
	require.NotEmpty(t, pixKey.Id)
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, descriptionTransaction)
	require.Empty(t, transaction.CancelDescription)

	_, err = model.NewTransaction(sourceAccount, amount, pixKey, descriptionTransaction)
	require.NotNil(t, err)

	_, err = model.NewTransaction(destinationAccount, 0, pixKey, descriptionTransaction)
	require.NotNil(t, err)
}

func TestModel_TransactionStatusChange(t *testing.T) {
	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	sourceAccountNumber := uuid.NewString()
	sourceOwnerName := "Moura"
	sourceAccount, _ := model.NewAccount(sourceOwnerName, sourceAccountNumber, bank)

	destinationAccountNumber := uuid.NewString()
	destinationOwnerName := "Moura"
	destinationAccount, _ := model.NewAccount(destinationOwnerName, destinationAccountNumber, bank)

	kind := "email"
	key := "email@email.com"
	pixKey, _ := model.NewPixKey(kind, key, sourceAccount)

	amount := 7.77
	descriptionTransaction := "My description"
	transaction1, _ := model.NewTransaction(destinationAccount, amount, pixKey, descriptionTransaction)

	transaction1.Confirm()
	require.Equal(t, transaction1.Status, model.TRANSACTION_STATUS_CONFIRMED)

	cancelDescription := "Error"
	transaction1.Cancel(cancelDescription)
	require.Equal(t, transaction1.Status, model.TRANSACTION_STATUS_CANCELLED)
	require.Equal(t, transaction1.CancelDescription, cancelDescription)

	transaction2, _ := model.NewTransaction(destinationAccount, amount, pixKey, descriptionTransaction)
	err := transaction2.Complete()
	require.NotNil(t, err)
}
