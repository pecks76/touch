package account

import accountDB "restservice/database/account"

//go:generate mockgen -source=pot_service.go -destination=mock/pot_service.go


type PotService interface {
	GetOrCreatePot(id int, name string, clientId int, depositId int) (int, error)
}

type potService struct {
	potRepository accountDB.PotRepository
}

func NewPotService(potRepository accountDB.PotRepository) PotService {
	return &potService{potRepository: potRepository}
}

func (ps potService) GetOrCreatePot(id int, name string, clientId int, depositId int) (int, error) {

	pot := ps.potRepository.ReadPot(id)
	if pot.Id == id {
		return id, nil
	}

	newId, err := ps.potRepository.InsertPot(name, clientId, depositId)
	if err != nil {
		return 0, err
	} else {
		return newId, nil
	}
}

