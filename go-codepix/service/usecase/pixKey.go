package usecase

import (
	"fmt"

	"github.com/moura1001/codepix/domain/model"
)

type PixKeyUseCase struct {
	repository model.IPixKeyRepository
}

func NewPixKeyUseCase(repository model.IPixKeyRepository) PixKeyUseCase {
	return PixKeyUseCase{
		repository: repository,
	}
}

func (p *PixKeyUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	errMsgTemplate := "error to register pix key in repository. Details: '%s'"

	account, err := p.repository.FindAccountById(accountId)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	pixKeyInserted, err := p.repository.RegisterKey(pixKey)
	if err != nil {
		return nil, fmt.Errorf(errMsgTemplate, err)
	}

	if pixKeyInserted == nil {
		return nil, fmt.Errorf(errMsgTemplate+"pix key (key=%s, kind=%s, accountId=%s) not inserted", key, kind, accountId)
	}

	if pixKeyInserted != nil || pixKeyInserted.Id != pixKey.Id {
		return nil, fmt.Errorf(errMsgTemplate+"an inconsistency was detected. The key entered in the database '%s' was different from the one generated by the server '%s'", pixKeyInserted.Id, pixKey.Id)
	}

	return pixKeyInserted, nil
}

func (p *PixKeyUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.repository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, fmt.Errorf("error to find pix key in repository. Details: '%s'", err)
	}
	return pixKey, nil
}