package createtransaction

import (
	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/masilvasql/fc-ms-wallet/internal/event"
	"github.com/masilvasql/fc-ms-wallet/internal/gateway"
	"github.com/masilvasql/fc-ms-wallet/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	CreateTransactionGateway gateway.TransactionGateway
	AccountGateway           gateway.AccountGateway
	Eventdispatcher          events.EventDispatcher
	TransactionCreated       events.EventInterface
}

func CreateNewTranasction(transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher *events.EventDispatcher,
	transactionCreated *event.TransactionCreated) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		CreateTransactionGateway: transactionGateway,
		AccountGateway:           accountGateway,
		Eventdispatcher:          *eventDispatcher,
		TransactionCreated:       transactionCreated,
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

	uc.TransactionCreated.SetPayload(output)
	uc.Eventdispatcher.Dispatch(uc.TransactionCreated)

	return output, nil

}
