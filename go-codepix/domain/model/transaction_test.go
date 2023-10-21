package model_test

import (
	"testing"

	"github.com/moura1001/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTransaction(t *testing.T) {
	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	sourceAccountNumber := "somenumber"
	sourceOwnerName := "Moura"
	sourceAccount, _ := model.NewAccount(sourceOwnerName, sourceAccountNumber, bank)

	destinationAccountNumber := "somenumber"
	destinationOwnerName := "Moura"
	destinationAccount, _ := model.NewAccount(destinationAccountNumber, destinationOwnerName, bank)

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
