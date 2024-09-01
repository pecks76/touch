package receipt

import (
	"github.com/golang/mock/gomock"
	mock "restservice/database/receipt/mock"
	"restservice/domain/receipt"
	"testing"
)

func TestDepositService_GetOrCreateDeposit_CreatesIfNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	depositRepo := mock.NewMockDepositRepository(ctrl)
	instructionRepo := mock.NewMockInstructionRepository(ctrl)

	service := NewDepositService(depositRepo, instructionRepo)

	depositRepo.EXPECT().ReadDeposit(123).Return(receipt.Deposit{Id: 0})
	depositRepo.EXPECT().InsertDeposit(456, 789)
	service.GetOrCreateDeposit(123, 456, 789)

}

func TestDepositService_GetOrCreateDeposit_DoesNotCreateIfFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	depositRepo := mock.NewMockDepositRepository(ctrl)
	instructionRepo := mock.NewMockInstructionRepository(ctrl)

	service := NewDepositService(depositRepo, instructionRepo)

	depositRepo.EXPECT().ReadDeposit(123).Return(receipt.Deposit{Id: 123})
	service.GetOrCreateDeposit(123, 456, 789)
}