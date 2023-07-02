package createaccount

import (
	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/masilvasql/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateay gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateay,
	}
}

func (uc *CreateAccountUseCase) Execute(inputDTO *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(inputDTO.ClientID)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	outputDTO := &CreateAccountOutputDTO{
		ID: account.ID,
	}

	return outputDTO, nil

}
