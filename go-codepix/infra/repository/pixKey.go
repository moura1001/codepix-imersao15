package repository

import (
	"errors"
	"fmt"

	"github.com/moura1001/codepix/domain/model"
	"gorm.io/gorm"
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
	err := r.Db.Create(&account).Error
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

func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (pixKey *model.PixKey, isErrRecordNotFound bool, err error) {
	var pix model.PixKey

	err = r.Db.Preload("Account.Bank").First(&pix, "kind=? and key=?", kind, key).Error
	pixKey = &pix
	isErrRecordNotFound = errors.Is(err, gorm.ErrRecordNotFound)

	if pix.Id == "" {
		pixKey = nil
		err = fmt.Errorf("no pix key was found for input params kind=%s and key=%s", kind, key)
	}

	return
}

func (r PixKeyRepositoryDb) FindAccountByNumber(bankCode string, number string) (account *model.Account, isErrRecordNotFound bool, err error) {
	var acc model.Account

	err = r.Db.Preload("Bank").First(&acc, "bank_code=? and number=?", bankCode, number).Error
	account = &acc
	isErrRecordNotFound = errors.Is(err, gorm.ErrRecordNotFound)

	if acc.Id == "" {
		account = nil
		err = fmt.Errorf("no account was found for input params bank_code=%s and number=%s", bankCode, number)
	}

	return
}

func (r PixKeyRepositoryDb) FindBankByCode(code string) (bank *model.Bank, isErrRecordNotFound bool, err error) {
	var ban model.Bank

	err = r.Db.First(&ban, "code=?", code).Error
	bank = &ban
	isErrRecordNotFound = errors.Is(err, gorm.ErrRecordNotFound)

	if ban.Id == "" {
		bank = nil
		err = fmt.Errorf("no bank was found for code=%s", code)
	}

	return
}
