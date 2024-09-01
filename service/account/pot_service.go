package account

import accountDB "restservice/database/account"

type PotService interface {
	GetOrCreatePot(id int, name string, clientId int, depositId int) int
}

type potService struct {
	potRepository accountDB.PotRepository
}

func NewPotService(potRepository accountDB.PotRepository) PotService {
	return &potService{potRepository: potRepository}
}

func (ps potService) GetOrCreatePot(id int, name string, clientId int, depositId int) int {

	pot := ps.potRepository.ReadPot(id)

	if pot.Id == 0 {
		return ps.potRepository.InsertPot(name, clientId, depositId)
	} else {
		return id
	}
}

