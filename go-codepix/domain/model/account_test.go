package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura1001/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	accountNumber := uuid.NewString()
	ownerName := "Moura"
	account, err := model.NewAccount(ownerName, accountNumber, bank)

	require.Nil(t, err)
	require.NotEmpty(t, account.Id)
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankCode, bank.Code)

	_, err = model.NewAccount(ownerName, "", bank)
	require.NotNil(t, err)
}
