package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var (
	DBConn *sql.DB
)

func init() {

	var err error
	// DBConn, err = sql.Open("mysql", "greg:greg@tcp(mysql:3306)/touch")
	DBConn, err = sql.Open("mysql", "greg:greg@tcp(localhost:3306)/touch")

	if err != nil {
		fmt.Println("Error in connecting to database")
		fmt.Printf("Error is: %+v\n ", err)
	}
}
