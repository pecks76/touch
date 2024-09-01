package account

import accountDb "restservice/database/account"

type ClientService interface {
	GetOrCreateClient(id int) int
}

type clientService struct {
	clientRepo accountDb.ClientRepository
}

func NewClientService(clientRepo accountDb.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

func (cs clientService) GetOrCreateClient(id int) int {

	client := cs.clientRepo.ReadClient(id)

	if client.Id == 0 {
		return cs.clientRepo.InsertClient()
	} else {
		return id
	}
}
