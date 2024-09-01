package account

type Pot struct {
	Id int
	Name   string
	ClientId int
	DepositId int
	Accounts []Account
}
