package model_test

import (
	"testing"

	"github.com/moura1001/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewBank(t *testing.T) {
	code := "001"
	name := "Caixa"

	bank, err := model.NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, bank.Id)
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", "")
	require.NotNil(t, err)
}
