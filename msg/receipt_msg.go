package msg

import "restservice/domain/account"

// not full domain objects, only for JSON / data transfer
type Account struct {
	Id          int
	WrapperType string
	PotId       int
	Amount      int
}
type Pot struct {
	Id       int
	Name     string
	ClientId int
	Accounts []Account
}

type Receipt struct {
	ID        int
	DepositId int
	Nominal   int
	ClientId  int
	Pots      []Pot
}

func PotFromDomainObj(from account.Pot) Pot {

	var to = Pot{
		Id:       from.Id,
		Name:     from.Name,
		ClientId: from.ClientId,
	}

	var accounts []Account
	// build up accounts as well
	for _, domAccount := range from.Accounts {
		msgAccount := accountFromDomainObj(domAccount)
		accounts = append(accounts, msgAccount)
	}

	to.Accounts = accounts

	return to
}

func accountFromDomainObj(from account.Account) Account {

	var to = Account{
		Id:          from.Id,
		WrapperType: from.WrapperType,
		PotId:       from.PotId,
		Amount:      from.Amount,
	}

	return to
}
