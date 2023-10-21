package model_test

import (
	"testing"

	"github.com/moura1001/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Caixa"
	bank, _ := model.NewBank(code, name)

	accountNumber := "somenumber"
	ownerName := "Moura"
	account, _ := model.NewAccount(ownerName, accountNumber, bank)

	kind := "email"
	key := "email@email.com"
	pixKey, err := model.NewPixKey(kind, key, account)

	require.Nil(t, err)
	require.NotEmpty(t, pixKey.Id)
	require.Equal(t, pixKey.Kind, model.PixKeyKind(kind))
	require.Equal(t, pixKey.Status, model.PIX_KEY_STATUS_ACTIVE)

	kind = "cpf"
	_, err = model.NewPixKey(kind, key, account)
	require.Nil(t, err)

	kind = "nome"
	_, err = model.NewPixKey(kind, key, account)
	require.NotNil(t, err)
}
