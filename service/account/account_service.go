package account

import (
	"errors"
	accountDb "restservice/database/account"
	"restservice/domain/account"
)

//go:generate mockgen -source=account_service.go -destination=mock/account_service.go

// all methods
type AccountService interface {
	AddToOrCreateAccount(id int, wrapperType string, potId int, amount int) error
}

// all dependencies
type accountService struct {
	accountRepo accountDb.AccountRepository
}

// constructor takes dependency and returns concrete impl, as Interface
func NewAccountService(repo accountDb.AccountRepository) AccountService {
	return &accountService{accountRepo: repo}
}

const SIPP = "SIPP"
const ISA = "ISA"
const GIA = "GIA"

func (ac accountService) AddToOrCreateAccount(id int, wrapperType string, potId int, amount int) error {

	accountTypes := []string{SIPP, ISA, GIA}
	if !contains(accountTypes, wrapperType) {
		return errors.New("invalid account type")
	}

	acc := ac.accountRepo.ReadAccount(id)

	if checkAmount(acc, amount, wrapperType) != nil {
		return errors.New("too much money in account")
	}

	var err error
	if acc.Id == 0 {
		err = ac.accountRepo.InsertAccount(wrapperType, potId, amount)
	} else {
		err = ac.accountRepo.UpdateAccount(acc.Id, amount)
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func checkAmount(account account.Account, amount int, wrapperType string) error {

	var total = account.Amount + amount

	if wrapperType == SIPP && total > 60000 {
		return errors.New("SIPP account can't hold more than £60000")
	}

	if wrapperType == ISA && total > 20000 {
		return errors.New("ISA account can't hold more than £20000")
	}
	return nil
}
