package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull" gorm:"column:owner_name;type:varchar(128);not null"`
	Number    string    `json:"number" valid:"required" gorm:"type:varchar(16)"`
	Bank      *Bank     `valid:"required"`
	BankId    string    `json:"bank_id" valid:"notnull" gorm:"column:bank_id;type:uuid;not null"`
	PixKeys   []*PixKey `valid:"-" gorm:"ForeignKey:AccountId"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(ownerName, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Number:    number,
		Bank:      bank,
	}
	if bank != nil {
		account.BankId = bank.Id
	}

	account.Id = uuid.NewString()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to create new account. Invalid input. Details: '%v'", err)
	}

	return &account, nil
}
