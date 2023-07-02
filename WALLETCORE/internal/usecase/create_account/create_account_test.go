package createaccount

import (
	"testing"

	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "teste@teste.com")
	mClient := &ClientGatewayMock{}
	mClient.On("Get", client.ID).Return(client, nil)

	mAccount := &AccountMock{}
	mAccount.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(mAccount, mClient)

	input := &CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	mClient.AssertExpectations(t)
	mClient.AssertNumberOfCalls(t, "Get", 1)
	mAccount.AssertExpectations(t)
	mAccount.AssertNumberOfCalls(t, "Save", 1)

}
