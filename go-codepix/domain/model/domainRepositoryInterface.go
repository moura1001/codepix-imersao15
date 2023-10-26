package model

type IPixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, bool, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccountByNumber(bankCode string, number string) (*Account, bool, error)
	FindBankByCode(code string) (*Bank, bool, error)
}

type ITransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	FindById(id string) (*Transaction, error)
}
