package account

import accountDb "restservice/database/account"

//go:generate mockgen -source=client_service.go -destination=mock/client_service.go

type ClientService interface {
	GetOrCreateClient(id int) (int, error)
}

type clientService struct {
	clientRepo accountDb.ClientRepository
}

func NewClientService(clientRepo accountDb.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

func (cs clientService) GetOrCreateClient(id int) (int, error) {

	client := cs.clientRepo.ReadClient(id)
	if client.Id == id {
		return id, nil
	}

	// else create
	newId, err := cs.clientRepo.InsertClient()
	if err != nil {
		return 0, err}
	return newId, nil

}
