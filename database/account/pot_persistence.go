package account

import (
	"log"
	"restservice/database"
	"restservice/domain/account"
)

type PotRepository interface {
	ReadPot(id int) account.Pot
	ReadPotsForDepositId(depositId int) []account.Pot
	InsertPot(name string, clientId int, depositId int) int
}

type potRepository struct {
	accountRepo AccountRepository
}

func NewPotRepository(accountRepo AccountRepository ) PotRepository {
	return &potRepository{accountRepo: accountRepo}
}

func (pr potRepository) ReadPot(id int) account.Pot {

	var pot account.Pot

	err := database.DBConn.QueryRow("SELECT id, name, clientId, depositId "+
		"FROM pot "+
		"WHERE id = ?", id).Scan(&pot.Id, &pot.Name, &pot.ClientId, &pot.DepositId)

	if err != nil {
		log.Println("Failed to read pot")
		// what to do at this point?
	}

	accounts := pr.accountRepo.ReadAccountsByPotId(id)
	pot.Accounts = accounts

	return pot

}

func (pr potRepository) ReadPotsForDepositId(depositId int) []account.Pot {

	rows, err := database.DBConn.Query("SELECT id, name, clientId, depositId FROM pot WHERE depositId = ?", depositId)

	if err != nil {
		log.Println("Failed to get pots by deposit id ")
		log.Printf("%+v", err)
	}

	var pots []account.Pot
	for rows.Next() {
		var pot account.Pot
		rows.Scan(&pot.Id, &pot.Name, &pot.ClientId, &pot.DepositId)

		accounts := pr.accountRepo.ReadAccountsByPotId(pot.Id)
		pot.Accounts = accounts

		pots = append(pots, pot)
	}

	return pots

}

func (pr potRepository) InsertPot(name string, clientId int, depositId int) int {

	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO pot(name, clientId, depositId) VALUES (?,?,?)")
	res, err := stmt.Exec(name, clientId, depositId)

	if err != nil {
		// todo: logging
		log.Printf("Failed to create pot: %+v", err)
	}

	id, _ := res.LastInsertId()
	return int(id)
}
