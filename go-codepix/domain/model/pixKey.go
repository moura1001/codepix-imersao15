package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

/*
*

	Pix Key Kind

*
*/
type PixKeyKind string

const (
	PIX_KEY_KIND_EMAIL PixKeyKind = "email"
	PIX_KEY_KIND_CPF   PixKeyKind = "cpf"
)

func (pKind PixKeyKind) isValidKind() bool {
	switch pKind {
	case PIX_KEY_KIND_EMAIL,
		PIX_KEY_KIND_CPF:
		return true
	}

	return false
}

func (pKind PixKeyKind) isValidValue(value string) bool {
	switch pKind {
	case PIX_KEY_KIND_EMAIL:
		return IsValideEmail(value)
	case PIX_KEY_KIND_CPF:
		return IsValideCpf(value)
	}

	return false
}

/*
*

	Pix Key Status

*
*/
type PixKeyStatus string

const (
	PIX_KEY_STATUS_ACTIVE   PixKeyStatus = "active"
	PIX_KEY_STATUS_INACTIVE PixKeyStatus = "inactive"
)

func (pStatus PixKeyStatus) isValid() bool {
	switch pStatus {
	case PIX_KEY_STATUS_ACTIVE,
		PIX_KEY_STATUS_INACTIVE:
		return true
	}

	return false
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      PixKeyKind   `json:"kind" valid:"notnull"`
	Key       string       `json:"key" valid:"notnull"`
	Account   *Account     `valid:"-"`
	AccountId string       `json:"account_id" valid:"notnull"`
	Status    PixKeyStatus `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	isValid := pixKey.Kind.isValidKind()
	if !isValid {
		return fmt.Errorf("invalid pix key kind type: %v", pixKey.Kind)
	}

	isValid = pixKey.Kind.isValidValue(pixKey.Key)
	if !isValid {
		return fmt.Errorf("invalid pix key value %s for type %v", pixKey.Key, pixKey.Kind)
	}

	isValid = pixKey.Status.isValid()
	if !isValid {
		return fmt.Errorf("invalid pix key status type: %v", pixKey.Status)
	}

	return nil
}

func NewPixKey(kind, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    PixKeyKind(kind),
		Key:     key,
		Account: account,
		Status:  PIX_KEY_STATUS_ACTIVE,
	}
	if account != nil {
		pixKey.AccountId = account.Id
	}

	pixKey.Id = uuid.NewString()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to create new pix key. Invalid input. Details: '%v'", err)
	}

	return &pixKey, nil
}
