package receipt

import (
	"errors"
	"log"
	"restservice/database"
	"restservice/domain/receipt"
)

//go:generate mockgen -source=instruction_persistence.go -destination=mock/instruction_persistence.go

type InstructionRepository interface {
	InsertInstruction(instruction receipt.Instruction) error
	ReadInstructionsByDepositId(depositId int) []receipt.Instruction
}

type instructionRepository struct {

}

func NewInstructionRepository() InstructionRepository {
	return &instructionRepository{}
}

func (ir instructionRepository) InsertInstruction(instruction receipt.Instruction) error {
	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO instruction(depositId, potName, wrapper_type, amount) VALUES (?,?,?,?)")
	_, err := stmt.Exec(instruction.DepositId, instruction.PotName, instruction.WrapperType, instruction.Amount)

	if err != nil {
		errors.New("failed to create instruction")
	}

	// id, _ := res.LastInsertId()
	return nil
}

func (ir instructionRepository) ReadInstructionsByDepositId(depositId int) []receipt.Instruction {

	rows, err := database.DBConn.Query("SELECT depositId, potName, wrapper_type, amount FROM instruction WHERE depositId = ?", depositId)
	if err != nil {
		// todo: should not fail, error here
		log.Println("Failed to get instructions by deposit id ")
	}

	var instructions []receipt.Instruction
	for rows.Next() {
		var instr receipt.Instruction
		rows.Scan(&instr.DepositId, &instr.PotName, &instr.WrapperType, &instr.Amount)
		instructions = append(instructions, instr)
	}

	return instructions
}