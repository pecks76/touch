package account

import (
	"github.com/golang/mock/gomock"
	"restservice/domain/account"
	"testing"
	//"github.com/stretchr/testify/assert"
	mock "restservice/database/account/mock"
)

func TestPotService_GetOrCreatePot_CreatesIfNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	potRepo := mock.NewMockPotRepository(ctrl)

	potRepo.EXPECT().ReadPot(123).Return(account.Pot{Id: 0})
	potRepo.EXPECT().InsertPot("A", 456, 789)

	service := NewPotService(potRepo)
	service.GetOrCreatePot(123, "A", 456, 789)
}

func TestPotService_GetOrCreatePot_DoesNotCreateIfFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	potRepo := mock.NewMockPotRepository(ctrl)

	potRepo.EXPECT().ReadPot(123).Return(account.Pot{Id: 123})

	service := NewPotService(potRepo)
	service.GetOrCreatePot(123, "A", 456, 789)
}