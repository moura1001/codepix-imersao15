package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

/*
*

	Transaction Status

*
*/
type TransactionStatus string

const (
	TRANSACTION_STATUS_PENDING   TransactionStatus = "pending"
	TRANSACTION_STATUS_CANCELLED TransactionStatus = "cancelled"
	TRANSACTION_STATUS_CONFIRMED TransactionStatus = "comfirmed"
	TRANSACTION_STATUS_COMPLETED TransactionStatus = "completed"
)

func (tStatus TransactionStatus) isValid() bool {
	switch tStatus {
	case TRANSACTION_STATUS_PENDING,
		TRANSACTION_STATUS_COMPLETED,
		TRANSACTION_STATUS_CANCELLED,
		TRANSACTION_STATUS_CONFIRMED:
		return true
	}

	return false
}

type Transaction struct {
	Base              `valid:"required"`
	AccountTo         *Account          `valid:"required" gorm:"foreignKey:AccountToNumber;references:Number"`
	AccountToNumber   string            `json:"account_to_number" valid:"notnull" gorm:"column:account_to_number;type:uuid;not null"`
	Amount            float64           `json:"amount" valid:"notnull" gorm:"type:float"`
	PixKeyFrom        *PixKey           `valid:"required"`
	PixKeyFromId      string            `json:"pix_key_from_id" valid:"notnull" gorm:"column:pix_key_from_id;type:uuid;not null"`
	Status            TransactionStatus `json:"status" valid:"notnull" gorm:"type:varchar(16)"`
	Description       string            `json:"description" valid:"-" gorm:"type:varchar(128)"`
	CancelDescription string            `json:"cancel_description" valid:"-" gorm:"type:varchar(128)"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}

	if transaction.PixKeyFrom.AccountNumber == transaction.AccountToNumber {
		return fmt.Errorf("the transaction source and destination account cannot be the same")
	}

	if transaction.Amount <= 0 {
		return fmt.Errorf("the transaction amount must be greater than 0")
	}

	isValid := transaction.Status.isValid()
	if !isValid {
		return fmt.Errorf("invalid transaction status type: %v", transaction.Status)
	}

	return nil
}

func NewTransaction(accountTo *Account, amount float64, pixKeyFrom *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountTo:   accountTo,
		Amount:      amount,
		PixKeyFrom:  pixKeyFrom,
		Status:      TRANSACTION_STATUS_PENDING,
		Description: description,
	}
	if accountTo != nil {
		transaction.AccountToNumber = accountTo.Number
	}
	if pixKeyFrom != nil {
		transaction.PixKeyFromId = pixKeyFrom.Id
	}

	transaction.Id = uuid.NewString()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to create new transaction. Invalid input. Details: '%v'", err)
	}

	return &transaction, nil
}

func (transaction *Transaction) Cancel(description string) error {
	if transaction.Status != TRANSACTION_STATUS_PENDING && transaction.Status != TRANSACTION_STATUS_CONFIRMED {
		return fmt.Errorf("invalid transaction status state change: %v -> %v", transaction.Status, TRANSACTION_STATUS_CANCELLED)
	}

	transaction.Status = TRANSACTION_STATUS_CANCELLED
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description
	return transaction.isValid()
}

func (transaction *Transaction) Confirm() error {
	if transaction.Status != TRANSACTION_STATUS_PENDING {
		return fmt.Errorf("invalid transaction status state change: %v -> %v", transaction.Status, TRANSACTION_STATUS_CONFIRMED)
	}

	transaction.Status = TRANSACTION_STATUS_CONFIRMED
	transaction.UpdatedAt = time.Now()
	return transaction.isValid()
}

func (transaction *Transaction) Complete() error {
	if transaction.Status != TRANSACTION_STATUS_CONFIRMED {
		return fmt.Errorf("invalid transaction status state change: %v -> %v", transaction.Status, TRANSACTION_STATUS_COMPLETED)
	}

	transaction.Status = TRANSACTION_STATUS_COMPLETED
	transaction.UpdatedAt = time.Now()
	return transaction.isValid()
}
