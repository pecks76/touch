package msg

import (
	domain "restservice/domain/receipt"
)

type Deposit struct {
	Id           int
	ClientId     int
	Nominal      int
	Pots         []Pot
	Instructions []Instruction
}

type Instruction struct {
	DepositId   int
	PotName     string
	WrapperType string
	Amount      int
}

func DepositFromDomainObj(from domain.Deposit, domainInstructions []domain.Instruction) Deposit {

	var to = Deposit{
		Id:       from.Id,
		ClientId: from.ClientId,
		Nominal:  from.Nominal,
	}

	var pots []Pot
	for _, p := range from.Pots {
		pot := PotFromDomainObj(p)
		pots = append(pots, pot)
	}

	to.Pots = pots

	var instructions []Instruction
	for _, i := range domainInstructions {
		instruction := Instruction{
			DepositId:   i.DepositId,
			PotName:     i.PotName,
			WrapperType: i.WrapperType,
			Amount:      i.Amount,
		}
		instructions = append(instructions, instruction)
	}

	to.Instructions = instructions

	return to
}
