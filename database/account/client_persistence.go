package account

import (
	"log"
	"restservice/database"
	"restservice/domain/account"
)

//go:generate mockgen -source=client_persistence.go -destination=mock/client_persistence.go

type ClientRepository interface {
	ReadClient(id int) account.Client
	InsertClient() int
}

type clientRepository struct {

}

func NewClientRepository() ClientRepository{
	return &clientRepository{}
}

func (cr clientRepository) ReadClient(id int) account.Client {

	var client = account.Client{Id: 0}

	err := database.DBConn.QueryRow("SELECT id " +
		"FROM client " +
		"WHERE id = ?", id).Scan(&client.Id)

	if err != nil {
		log.Println("Failed to read client")
	}

	return client
}

func (cr clientRepository) InsertClient() int {

	// we need to use this pattern to get the Id back out
	stmt, _ := database.DBConn.Prepare("INSERT INTO client values()")
	res, err := stmt.Exec()

	if err != nil {
		log.Println("Failed to create client")
	}

	// todo: error checking here
	log.Printf("%+v", res)
	id, _ :=  res.LastInsertId()
	return int(id)
}
