package receipt

import (
	receiptDB "restservice/database/receipt"
	"restservice/domain/receipt"
)

type InstructionService interface {
	saveInstruction(depositId int, potName string, wrapperType string, amount int)
}

type instructionService struct {
	instructionRepo receiptDB.InstructionRepository
}

func NewInstructionService(instructionRepo receiptDB.InstructionRepository) InstructionService {
	return &instructionService{instructionRepo: instructionRepo}
}

func (is instructionService) saveInstruction(depositId int, potName string, wrapperType string, amount int) {

	var instruction = receipt.Instruction{
		DepositId:   depositId,
		PotName:     potName,
		WrapperType: wrapperType,
		Amount:      amount,
	}

	is.instructionRepo.InsertInstruction(instruction)
}

func (is instructionService) GetInstructionsByDepositId(depositId int) []receipt.Instruction {
	return is.instructionRepo.ReadInstructionsByDepositId(depositId)
}
