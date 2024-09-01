package receipt

import (
	"log"
	"restservice/database"
	"restservice/domain/receipt"
)

//go:generate mockgen -source=instruction_persistence.go -destination=mock/instruction_persistence.go

type InstructionRepository interface {
	InsertInstruction(instruction receipt.Instruction) int
	ReadInstructionsByDepositId(depositId int) []receipt.Instruction
}

type instructionRepository struct {

}

func NewInstructionRepository() InstructionRepository {
	return &instructionRepository{}
}

func (ir instructionRepository) InsertInstruction(instruction receipt.Instruction) int {
	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO instruction(depositId, potName, wrapper_type, amount) VALUES (?,?,?,?)")
	res, err := stmt.Exec(instruction.DepositId, instruction.PotName, instruction.WrapperType, instruction.Amount)

	if err != nil {
		// todo: logging
		log.Println("Failed to create instruction")
	}

	id, _ := res.LastInsertId()
	return int(id)
}

func (ir instructionRepository) ReadInstructionsByDepositId(depositId int) []receipt.Instruction {

	rows, err := database.DBConn.Query("SELECT depositId, potName, wrapper_type, amount FROM instruction WHERE depositId = ?", depositId)
	if err != nil {
		log.Println("Failed to get instructions by deposit id ")
	}

	var instructions []receipt.Instruction
	for rows.Next() {
		var instr receipt.Instruction
		rows.Scan(&instr.DepositId, &instr.PotName, &instr.WrapperType, &instr.Amount)
		instructions = append(instructions, instr)
	}

	// todo: there is not enough error checking going on here

	return instructions
}