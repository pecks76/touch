package receipt

import (
	"restservice/msg"
	accountService "restservice/service/account"
)

type ReceiptService interface {
	SaveReceipt(receipt msg.Receipt) error
}

type receiptService struct {
	clientService      accountService.ClientService
	depositService     DepositService
	potService         accountService.PotService
	accountService     accountService.AccountService
	instructionService InstructionService
}

func NewReceiptService(clientService accountService.ClientService,
	depositService DepositService,
	potService accountService.PotService,
	accountService accountService.AccountService,
	instructionService InstructionService) ReceiptService {
	return &receiptService{clientService: clientService,
		depositService:     depositService,
		potService:         potService,
		accountService:     accountService,
		instructionService: instructionService}
}

func (rc receiptService) SaveReceipt(receipt msg.Receipt) error {

	var clientId int
	var err error
	if clientId, err = rc.clientService.GetOrCreateClient(receipt.ClientId); err != nil {
		return err
	}

	var depositId int
	if depositId, err = rc.depositService.GetOrCreateDeposit(receipt.DepositId, clientId, receipt.Nominal); err != nil {
		return err
	}

	for _, p := range receipt.Pots {

		var potId int
		if potId, err = rc.potService.GetOrCreatePot(p.Id, p.Name, clientId, depositId); err != nil{
			return err
		}
		for _, a := range p.Accounts {
			if err = rc.accountService.AddToOrCreateAccount(a.Id, a.WrapperType, potId, a.Amount); err != nil {
				return err
			} else {
				return rc.instructionService.saveInstruction(depositId, p.Name, a.WrapperType, a.Amount)
			}
		}
	}

	return nil
}
