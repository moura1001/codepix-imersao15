package model

type IPixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccountById(id string) (*Account, error)
	FindBankById(id string) (*Bank, error)
}

type ITransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	FindById(id string) (*Transaction, error)
}
