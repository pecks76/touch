package account

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "restservice/database/account/mock"
	"restservice/domain/account"
	"testing"
)

func TestAccountService_AddToOrCreateAccount_ErrorOnInvalidWrapper(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "BADGER", 456, 789)

	expectedMsg := "invalid account type"
	assert.EqualErrorf(t, err, expectedMsg, "Error should be: %v, got: %v", expectedMsg, err)
}

func TestAccountService_AddToOrCreateAccount_Create_NoErrorOnValidSIPPAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 0})
	accountRepo.EXPECT().InsertAccount("SIPP", 456, 60000)

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "SIPP", 456, 60000)

	assert.Nil(t, err, "")
}

func TestAccountService_AddToOrCreateAccount_AddTo_NoErrorOnValidSIPPAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 123})
	accountRepo.EXPECT().UpdateAccount(123, 60000)

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "SIPP", 456, 60000)

	assert.Nil(t, err, "")
}

func TestAccountService_AddToOrCreateAccount_Create_ErrorOnInvalidSIPPAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 0})

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "SIPP", 456, 60001)

	expectedMsg := "too much money in account"
	assert.EqualErrorf(t, err, expectedMsg, "Error should be: %v, got: %v", expectedMsg, err)
}

func TestAccountService_AddToOrCreateAccount_AddTo_ErrorOnInvalidSIPPAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 123, Amount: 60000})

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "SIPP", 456, 1)

	expectedMsg := "too much money in account"
	assert.EqualErrorf(t, err, expectedMsg, "Error should be: %v, got: %v", expectedMsg, err)
}

func TestAccountService_AddToOrCreateAccount_NoErrorOnValidISAAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 0})
	accountRepo.EXPECT().InsertAccount("ISA", 456, 20000)

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "ISA", 456, 20000)

	assert.Nil(t, err, "")
}

func TestAccountService_AddToOrCreateAccount_Create_ErrorOnInvalidISAAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 0})

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "ISA", 456, 20001)

	expectedMsg := "too much money in account"
	assert.EqualErrorf(t, err, expectedMsg, "Error should be: %v, got: %v", expectedMsg, err)
}

func TestAccountService_AddToOrCreateAccount_AddTo_ErrorOnInvalidISAAmount(t *testing.T) {

	ctrl := gomock.NewController(t)
	accountRepo := mock.NewMockAccountRepository(ctrl)

	accountRepo.EXPECT().ReadAccount(123).Return(account.Account{Id: 123, Amount: 20000})

	service := NewAccountService(accountRepo)
	err := service.AddToOrCreateAccount(123, "ISA", 456, 1)

	expectedMsg := "too much money in account"
	assert.EqualErrorf(t, err, expectedMsg, "Error should be: %v, got: %v", expectedMsg, err)
}