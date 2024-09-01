package account

import (
	"errors"
	"log"
	accountDb "restservice/database/account"
	"restservice/domain/account"
)

// all methods
type AccountService interface {
	AddToOrCreateAccount(id int, wrapperType string, potId int, amount int)
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

func (ac accountService) AddToOrCreateAccount(id int, wrapperType string, potId int, amount int) {

	account := ac.accountRepo.ReadAccount(id)

	accountTypes := []string{SIPP, ISA, GIA}
	if !contains(accountTypes, wrapperType) {
		// todo: error handling here
		// invalid wrapperType
		// throw error
	}

	if checkAmount(account, amount, wrapperType) != nil {
		// todo: error - return error
		log.Println("Received error, too much money in account")
	}

	if account.Id == 0 {
		// check amount
		ac.accountRepo.InsertAccount(wrapperType, potId, amount)
	} else {
		// check amount (existing + instruction)
		ac.accountRepo.UpdateAccount(account.Id, amount)
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
