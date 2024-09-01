package receipt

import (
	"restservice/domain/account"
)

type Deposit struct {
	Id       int
	ClientId int
	Nominal  int
	Pots     []account.Pot
}
