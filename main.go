package main

import (
	"fmt"
	"log"
	"net/http"
	accountDB "restservice/database/account"
	receiptDB "restservice/database/receipt"
	"restservice/rest"
	"restservice/service/account"
	"restservice/service/receipt"
)

func main() {

	// wire dependencies through

	accountRepository := accountDB.NewAccountRepository()
	clientRepository := accountDB.NewClientRepository()
	instructionRepository := receiptDB.NewInstructionRepository()
	potRepository := accountDB.NewPotRepository(accountRepository)
	depositRepository := receiptDB.NewDepositRepository(potRepository)

	accountService := account.NewAccountService(accountRepository)
	clientService := account.NewClientService(clientRepository)
	instructionService := receipt.NewInstructionService(instructionRepository)
	potService := account.NewPotService(potRepository)
	depositoryService := receipt.NewDepositService(depositRepository, instructionRepository)
	depositService  := receipt.NewDepositService(depositRepository, instructionRepository)
	receiptService := receipt.NewReceiptService(clientService, depositoryService, potService, accountService, instructionService )

	receiptRestService := rest.NewReceiptRestService(receiptService)
	depositRestService := rest.NewDepositRestService(depositService)

	http.HandleFunc("/deposit/{id}", depositRestService.GET)
	http.HandleFunc("/receipt", receiptRestService.POST)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

