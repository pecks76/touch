package receipt

import (
	"restservice/msg"
	accountService "restservice/service/account"
)

type ReceiptService interface {
	SaveReceipt(receipt msg.Receipt)
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

func (rc receiptService) SaveReceipt(receipt msg.Receipt) {

	clientId := rc.clientService.GetOrCreateClient(receipt.ClientId)
	depositId := rc.depositService.GetOrCreateDeposit(receipt.DepositId, clientId, receipt.Nominal)

	for _, p := range receipt.Pots {
		potId := rc.potService.GetOrCreatePot(p.Id, p.Name, clientId, depositId)
		for _, a := range p.Accounts {
			rc.accountService.AddToOrCreateAccount(a.Id, a.WrapperType, potId, a.Amount)
			rc.instructionService.saveInstruction(depositId, p.Name, a.WrapperType, a.Amount)
		}
	}
}
