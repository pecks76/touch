package receipt

import (
	"errors"
	"log"
	"restservice/database"
	"restservice/database/account"
	"restservice/domain/receipt"
)

//go:generate mockgen -source=deposit_persistence.go -destination=mock/desposit_persistence.go

type DepositRepository interface {
	InsertDeposit(clientId int, nominal int) (int, error)
	ReadDeposit(id int) receipt.Deposit
}

type depositRepository struct {
	potRepo account.PotRepository
}

func NewDepositRepository(potRepo account.PotRepository) DepositRepository {
	return &depositRepository{potRepo: potRepo}
}

func (dr depositRepository) InsertDeposit(clientId int, nominal int) (int, error) {
	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO deposit(clientId, nominal) VALUES (?,?)")
	res, err := stmt.Exec(clientId, nominal)

	if err != nil {
		return 0, errors.New("failed to create deposit")
	}

	id, _ := res.LastInsertId()
	return int(id), nil
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

	pots, err := dr.potRepo.ReadPotsForDepositId(deposit.Id)
	if err != nil {
		log.Println(err)
	}
	deposit.Pots = pots
	log.Printf("deposit.Pots: %+v", deposit.Pots)

	return deposit
}