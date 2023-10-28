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
	Number    string    `json:"number" valid:"uuid" gorm:"primaryKey;type:uuid"`
	Bank      *Bank     `valid:"required" gorm:"foreignKey:BankCode;references:Code"`
	BankCode  string    `json:"bank_code" valid:"notnull" gorm:"column:bank_code;type:varchar(8)"`
	PixKeys   []*PixKey `valid:"-" gorm:"foreignKey:AccountNumber"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(ownerName string, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Number:    number,
		Bank:      bank,
	}
	if bank != nil {
		account.BankCode = bank.Code
	}

	account.Id = uuid.NewString()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to create new account. Invalid input. Details: '%v'", err)
	}

	return &account, nil
}
