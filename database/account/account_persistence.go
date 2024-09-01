package account

import (
	"log"
	"restservice/database"
	"restservice/domain/account"
)

//go:generate mockgen -source=account_persistence.go -destination=mock/account_persistence.go

type AccountRepository interface {
	ReadAccount(id int) account.Account
	ReadAccountsByPotId(potId int) []account.Account
	InsertAccount(wrapperType string, potId int, amount int) int
	UpdateAccount(id int, amount int)
}

type accountRepository struct {

}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (ar accountRepository) ReadAccount(id int) account.Account {

	var account account.Account

	err := database.DBConn.QueryRow("SELECT id, wrapper_type, potId, amount "+
		"FROM account "+
		"WHERE id = ?", id).Scan(&account.Id, &account.WrapperType, &account.PotId, &account.Amount)

	if err != nil {
		// todo: logging
		log.Println("Failed to read account")
	}

	return account
}

func (ar accountRepository) ReadAccountsByPotId(potId int) []account.Account {

	rows, err := database.DBConn.Query("SELECT id, wrapper_type, potId, amount FROM account WHERE potID = ?", potId)
	if err != nil {
		log.Println("Failed to get accounts by pot id ")
	}

	var accounts []account.Account
	for rows.Next() {
		var acc account.Account
		rows.Scan(&acc.Id, &acc.WrapperType, &acc.PotId, &acc.Amount)
		accounts = append(accounts, acc)
	}

	// todo: there is not enough error checking going on here

	return accounts
}

func (ar accountRepository) InsertAccount(wrapperType string, potId int, amount int) int {

	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO account(wrapper_type, potId, amount) VALUES (?,?,?)")
	res, err := stmt.Exec(wrapperType, potId, amount)

	if err != nil {
		// todo: this is definitely an error
		log.Println("Failed to create account")
	}

	id, _ := res.LastInsertId()
	return int(id)
}

func (ar accountRepository) UpdateAccount(id int, amount int) {

	_, err := database.DBConn.Query("UPDATE account SET amount = ? WHERE id = ?", amount, id)

	if err != nil {
		// todo: this is definitely an error
		log.Println("Failed to update account")
	}
}
