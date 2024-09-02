package receipt

import (
	receiptDB "restservice/database/receipt"
	"restservice/domain/receipt"
)

//go:generate mockgen -source=instruction_service.go -destination=mock/instruction_service.go

type InstructionService interface {
	saveInstruction(depositId int, potName string, wrapperType string, amount int) error
}

type instructionService struct {
	instructionRepo receiptDB.InstructionRepository
}

func NewInstructionService(instructionRepo receiptDB.InstructionRepository) InstructionService {
	return &instructionService{instructionRepo: instructionRepo}
}

func (is instructionService) saveInstruction(depositId int, potName string, wrapperType string, amount int) error {

	var instruction = receipt.Instruction{
		DepositId:   depositId,
		PotName:     potName,
		WrapperType: wrapperType,
		Amount:      amount,
	}

	return is.instructionRepo.InsertInstruction(instruction)
}
