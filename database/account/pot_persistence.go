package account

import (
	"errors"
	"log"
	"restservice/database"
	"restservice/domain/account"
)

//go:generate mockgen -source=pot_persistence.go -destination=mock/pot_persistence.go

type PotRepository interface {
	ReadPot(id int) account.Pot
	ReadPotsForDepositId(depositId int) ([]account.Pot, error)
	InsertPot(name string, clientId int, depositId int) (int, error)
}

type potRepository struct {
	accountRepo AccountRepository
}

func NewPotRepository(accountRepo AccountRepository ) PotRepository {
	return &potRepository{accountRepo: accountRepo}
}

func (pr potRepository) ReadPot(id int) account.Pot{

	var pot account.Pot

	err := database.DBConn.QueryRow("SELECT id, name, clientId, depositId "+
		"FROM pot "+
		"WHERE id = ?", id).Scan(&pot.Id, &pot.Name, &pot.ClientId, &pot.DepositId)

	if err != nil {
		log.Println("Failed to read pot")
		return pot
	}

	accounts, err := pr.accountRepo.ReadAccountsByPotId(id)
	if err != nil {
		log.Println(err)
	}
	pot.Accounts = accounts

	return pot

}

func (pr potRepository) ReadPotsForDepositId(depositId int) ([]account.Pot, error) {

	rows, err := database.DBConn.Query("SELECT id, name, clientId, depositId FROM pot WHERE depositId = ?", depositId)

	if err != nil {
		return nil, errors.New("failed to get pots by deposit id")
	}

	var pots []account.Pot
	for rows.Next() {
		var pot account.Pot
		rows.Scan(&pot.Id, &pot.Name, &pot.ClientId, &pot.DepositId)

		accounts, err := pr.accountRepo.ReadAccountsByPotId(pot.Id)
		if err != nil {
			return nil, errors.New("failed to get accounts by pot id")
		}
		pot.Accounts = accounts

		pots = append(pots, pot)
	}

	return pots, nil

}

func (pr potRepository) InsertPot(name string, clientId int, depositId int) (int, error) {

	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO pot(name, clientId, depositId) VALUES (?,?,?)")
	res, err := stmt.Exec(name, clientId, depositId)

	if err != nil {
		return 0, errors.New("failed to create pot")
	}

	id, _ := res.LastInsertId()
	return int(id), nil
}
