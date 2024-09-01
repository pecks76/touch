package account

import (
	"github.com/golang/mock/gomock"
	"restservice/domain/account"
	"testing"
	//"github.com/stretchr/testify/assert"
	mock "restservice/database/account/mock"
)

func TestClientService_GetOrCreateClient_CreatesIfNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	clientRepo := mock.NewMockClientRepository(ctrl)

	// if Read returns id 0
	clientRepo.EXPECT().ReadClient(123).Return(account.Client{Id: 0})
	clientRepo.EXPECT().InsertClient()

	service := NewClientService(clientRepo)
	service.GetOrCreateClient(123)
}

func TestClientService_GetOrCreateClient_DoesNotCreateIfFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	clientRepo := mock.NewMockClientRepository(ctrl)

	// if Read returns id 0
	clientRepo.EXPECT().ReadClient(123).Return(account.Client{Id: 123})

	service := NewClientService(clientRepo)
	service.GetOrCreateClient(123)
}
