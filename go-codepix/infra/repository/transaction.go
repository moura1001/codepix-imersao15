package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/moura1001/codepix/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return fmt.Errorf("error to insert transaction on database. Details: '%s'", err)
	}
	return nil
}

func (r TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return fmt.Errorf("error to update transaction on database. Details: '%s'", err)
	}
	return nil
}

func (r TransactionRepositoryDb) FindById(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountTo.Bank").Find(&transaction, "id=?", id)

	if transaction.Id == "" {
		return nil, fmt.Errorf("no transaction was found for id=%s", id)
	}

	return &transaction, nil
}
