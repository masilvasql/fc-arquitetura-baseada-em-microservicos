package createtransaction

import (
	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/masilvasql/fc-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	CreateTransactionGateway gateway.TransactionGateway
	AccountGateway           gateway.AccountGateway
}

func CreateNewTranasction(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		CreateTransactionGateway: transactionGateway,
		AccountGateway:           accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {

	accFrom, err := uc.AccountGateway.FindById(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accTo, err := uc.AccountGateway.FindById(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accFrom, accTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.CreateTransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{
		TransactionID: transaction.ID,
	}

	return output, nil

}
