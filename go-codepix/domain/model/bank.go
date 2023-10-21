package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.Id = uuid.NewString()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to create new bank. Invalid input. Details: '%v'", err)
	}

	return &bank, nil
}
