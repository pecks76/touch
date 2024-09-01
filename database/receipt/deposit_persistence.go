package receipt

import (
	"log"
	"restservice/database"
	"restservice/database/account"
	"restservice/domain/receipt"
)

type DepositRepository interface {
	InsertDeposit(clientId int, nominal int) int
	ReadDeposit(id int) receipt.Deposit
}

type depositRepository struct {
	potRepo account.PotRepository
}

func NewDepositRepository(potRepo account.PotRepository) DepositRepository {
	return &depositRepository{potRepo: potRepo}
}

func (dr depositRepository) InsertDeposit(clientId int, nominal int) int {
	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO deposit(clientId, nominal) VALUES (?,?)")
	res, err := stmt.Exec(clientId, nominal)

	if err != nil {
		// todo: logging
		log.Println("Failed to create deposit")
	}

	id, _ := res.LastInsertId()
	return int(id)
}


func (dr depositRepository) ReadDeposit(id int) receipt.Deposit {

	var deposit = receipt.Deposit{Id: 0}
	var err = database.DBConn.QueryRow(
		`SELECT id, clientID, nominal
		FROM deposit d
		WHERE id = ?`, id).Scan(&deposit.Id, &deposit.ClientId, &deposit.Nominal)

	if err != nil {
		log.Println("Failed to read deposit")
	}

	pots := dr.potRepo.ReadPotsForDepositId(deposit.Id)
	deposit.Pots = pots
	log.Printf("deposit.Pots: %+v", deposit.Pots)

	return deposit
}