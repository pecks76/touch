package receipt

import (
	receiptDB "restservice/database/receipt"
	"restservice/domain/receipt"
	"restservice/msg"
)

type DepositService interface {
	GetDepositReport(id int) msg.Deposit
	GetOrCreateDeposit(id int, clientId int, nominal int) int
}

type depositService struct {
	depositRepo     receiptDB.DepositRepository
	instructionRepo receiptDB.InstructionRepository
}

func NewDepositService(depositRepo receiptDB.DepositRepository,
	instructionRepo receiptDB.InstructionRepository) DepositService {
	return &depositService{depositRepo: depositRepo,
		instructionRepo: instructionRepo}
}

func (ds depositService) GetDepositReport(id int) msg.Deposit {

	deposit := ds.depositRepo.ReadDeposit(id)
	var instructions []receipt.Instruction
	if deposit.Id != 0 {
		instructions = ds.instructionRepo.ReadInstructionsByDepositId(deposit.Id)
	}
	depositMsg := msg.DepositFromDomainObj(deposit, instructions)
	return depositMsg
}

func (ds depositService) GetOrCreateDeposit(id int, clientId int, nominal int) int {

	deposit := ds.depositRepo.ReadDeposit(id)

	if deposit.Id != 0 {
		return id
	} else {
		return ds.depositRepo.InsertDeposit(clientId, nominal)
	}
}
