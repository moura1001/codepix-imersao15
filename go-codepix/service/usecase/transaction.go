package usecase

import (
	"fmt"

	"github.com/moura1001/codepix/domain/model"
)

type TransactionUseCase struct {
	transactionRepository model.ITransactionRepository
	pixRepository         model.IPixKeyRepository
}

func NewTransactionUseCase(transactionRepository model.ITransactionRepository, pixRepository model.IPixKeyRepository) TransactionUseCase {
	return TransactionUseCase{
		transactionRepository: transactionRepository,
		pixRepository:         pixRepository,
	}
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyFrom, pixKeyFromKind, description string) (*model.Transaction, error) {
	errMsgTemplate := "error to register transaction in repository. Details: '%s'"

	account, err := t.pixRepository.FindAccountById(accountId)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	pixKey, err := t.pixRepository.FindKeyByKind(pixKeyFrom, pixKeyFromKind)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	return transaction, nil
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	errMsgTemplate := "error to confirm transaction. Details: '%s'"

	transaction, err := t.transactionRepository.FindById(transactionId)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = transaction.Confirm()
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	errMsgTemplate := "error to complete transaction. Details: '%s'"

	transaction, err := t.transactionRepository.FindById(transactionId)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = transaction.Complete()
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	return transaction, nil
}

func (t *TransactionUseCase) Cancel(transactionId string, reason string) (*model.Transaction, error) {
	errMsgTemplate := "error to cancel transaction. Details: '%s'"

	transaction, err := t.transactionRepository.FindById(transactionId)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = transaction.Cancel(reason)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	err = t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	return transaction, nil
}
