package model

type IPixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKing(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccountById(id string) (*Account, error)
}

type ITransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	FindById(id string) (*Transaction, error)
}
