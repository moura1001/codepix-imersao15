package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/moura1001/codepix/domain/model"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return fmt.Errorf("error to insert bank on database. Details: '%s'", err)
	}
	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return fmt.Errorf("error to insert account on database. Details: '%s'", err)
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, fmt.Errorf("error to insert pix key on database. Details: '%s'", err)
	}
	return pixKey, nil
}

func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").Find(&pixKey, "kind=? and key=?", kind, key)

	if pixKey.Id == "" {
		return nil, fmt.Errorf("no pix key was found for input params kind=%s and key=%s", kind, key)
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindAccountById(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").Find(&account, "id=?", id)

	if account.Id == "" {
		return nil, fmt.Errorf("no account was found for id=%s", id)
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) FindBankById(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.Find(&bank, "id=?", id)

	if bank.Id == "" {
		return nil, fmt.Errorf("no bank was found for id=%s", id)
	}

	return &bank, nil
}
