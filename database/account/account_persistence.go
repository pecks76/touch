package account

import (
	"errors"
	"log"
	"restservice/database"
	"restservice/domain/account"
)

//go:generate mockgen -source=account_persistence.go -destination=mock/account_persistence.go

type AccountRepository interface {
	ReadAccount(id int) account.Account
	ReadAccountsByPotId(potId int) ([]account.Account, error)
	InsertAccount(wrapperType string, potId int, amount int) error
	UpdateAccount(id int, amount int) error
}

type accountRepository struct {

}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (ar accountRepository) ReadAccount(id int) account.Account {

	var acc = account.Account{Id: 0}

	err := database.DBConn.QueryRow("SELECT id, wrapper_type, potId, amount "+
		"FROM account "+
		"WHERE id = ?", id).Scan(&acc.Id, &acc.WrapperType, &acc.PotId, &acc.Amount)

	if err != nil {
		log.Println("Failed to read account")
	}

	return acc
}

func (ar accountRepository) ReadAccountsByPotId(potId int) ([]account.Account, error) {

	rows, err := database.DBConn.Query("SELECT id, wrapper_type, potId, amount FROM account WHERE potID = ?", potId)

	if err != nil {
		return nil, errors.New("failed to get accounts by pot id")
	}

	var accounts []account.Account
	for rows.Next() {
		var acc account.Account
		rows.Scan(&acc.Id, &acc.WrapperType, &acc.PotId, &acc.Amount)
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (ar accountRepository) InsertAccount(wrapperType string, potId int, amount int) error {

	stmt, _ := database.DBConn.Prepare("INSERT INTO account(wrapper_type, potId, amount) VALUES (?,?,?)")
	_, err := stmt.Exec(wrapperType, potId, amount)

	if err != nil {
		return errors.New("failed to create account")
	}

	return nil
}

func (ar accountRepository) UpdateAccount(id int, amount int) error {

	_, err := database.DBConn.Query("UPDATE account SET amount = ? WHERE id = ?", amount, id)

	if err != nil {
		return errors.New("failed to update account")
	}

	return nil
}
