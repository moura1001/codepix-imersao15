package dto

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/moura1001/codepix/domain/model"
)

type TransactionDTOInputNew struct {
	RelatedTransactionIdFrom string  `json:"relatedTransactionIdFrom" valid:"required,uuid"`
	BankCodeFrom             string  `json:"bankCodeFrom" valid:"required"`
	BankCodeTo               string  `json:"bankCodeTo" valid:"required"`
	AccountNumberTo          string  `json:"accountNumberTo" valid:"required,uuid"`
	Amount                   float64 `json:"amount" valid:"required,numeric"`
	PixKeyFrom               string  `json:"pixKeyFrom" valid:"required"`
	PixKeyFromKind           string  `json:"pixKeyFromKind" valid:"required"`
	Description              string  `json:"description" valid:"-"`
}

func (transaction *TransactionDTOInputNew) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionDTOInputNew(data []byte) (*TransactionDTOInputNew, error) {
	var transaction TransactionDTOInputNew

	err := json.Unmarshal(data, &transaction)
	if err != nil {
		return nil, fmt.Errorf("error to parse new transaction for input data '%s'. Details: '%v'", string(data), err)
	}

	err = transaction.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to validate new transaction for input data '%s'. Details: '%v'", string(data), err)
	}

	return &transaction, nil
}

type TransactionDTOInputExistent struct {
	Id                     string                  `json:"id" valid:"required,uuid"`
	Status                 model.TransactionStatus `json:"status" valid:"required"`
	CancelDescription      string                  `json:"cancelDescription" valid:"-"`
	TransactionDTOInputNew `valid:"required"`
}

func (transaction *TransactionDTOInputExistent) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionDTOInputExistent(data []byte) (*TransactionDTOInputExistent, error) {
	var transaction TransactionDTOInputExistent

	err := json.Unmarshal(data, &transaction)
	if err != nil {
		return nil, fmt.Errorf("error to parse existent transaction for input data '%s'. Details: '%v'", string(data), err)
	}

	err = transaction.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to validate existent transaction for input data '%s'. Details: '%v'", string(data), err)
	}

	return &transaction, nil
}

type TransactionDTOOutput struct {
	TransactionDTOInputExistent `valid:"required"`
}

func (transaction *TransactionDTOOutput) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionDTOOutputJsonNew(id string, status model.TransactionStatus, errorMsg string, transactionInput TransactionDTOInputNew) ([]byte, error) {
	transactionOutput := TransactionDTOOutput{
		TransactionDTOInputExistent: TransactionDTOInputExistent{
			Id:                     id,
			Status:                 status,
			CancelDescription:      errorMsg,
			TransactionDTOInputNew: transactionInput,
		},
	}

	err := transactionOutput.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to validate new transaction json output data '%v'. Details: '%v'", transactionOutput, err)
	}

	data, err := json.Marshal(&transactionOutput)
	if err != nil {
		return nil, fmt.Errorf("error to encoding new transaction json output data '%v'. Details: '%v'", transactionOutput, err)
	}

	return data, nil
}

func NewTransactionDTOOutputJsonExistent(transactionInput TransactionDTOInputExistent) ([]byte, error) {
	transactionOutput := TransactionDTOOutput{
		TransactionDTOInputExistent: transactionInput,
	}

	err := transactionOutput.isValid()
	if err != nil {
		return nil, fmt.Errorf("error to validate existent transaction json output data '%v'. Details: '%v'", transactionOutput, err)
	}

	data, err := json.Marshal(&transactionOutput)
	if err != nil {
		return nil, fmt.Errorf("error to encoding existent transaction json output data '%v'. Details: '%v'", transactionOutput, err)
	}

	return data, nil
}
