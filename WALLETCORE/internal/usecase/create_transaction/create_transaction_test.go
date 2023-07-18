package createtransaction

import (
	"github.com/masilvasql/fc-ms-wallet/internal/event"
	"github.com/masilvasql/fc-ms-wallet/pkg/events"
	"testing"

	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountMock struct {
	mock.Mock
}

func (m *AccountMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionMock struct {
	mock.Mock
}

func (m *TransactionMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactiontUseCase_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("John Doe", "john@teste.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(200.0)

	clientTo, _ := entity.NewClient("Pedro", "pedro@teste.com")
	accountTo := entity.NewAccount(clientTo)
	ammount := 100.0

	mAccount := &AccountMock{}
	mAccount.On("FindById", accountFrom.ID).Return(accountFrom, nil)
	mAccount.On("FindById", accountTo.ID).Return(accountTo, nil)

	mTransaction := &TransactionMock{}
	mTransaction.On("Create", mock.Anything).Return(nil)

	input := &CreateTransactionInputDTO{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        ammount,
	}
	dispatcher := *events.NewEventDispatcher()
	event := event.NewTrasactionCreated()
	uc := CreateNewTranasction(mTransaction, mAccount, dispatcher, *event)

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mAccount.AssertExpectations(t)
	mAccount.AssertNumberOfCalls(t, "FindById", 2)
	mTransaction.AssertExpectations(t)
	mTransaction.AssertNumberOfCalls(t, "Create", 1)

}
