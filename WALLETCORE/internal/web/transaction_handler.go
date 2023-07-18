package web

import (
	"encoding/json"
	createtransaction "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_transaction"
	"net/http"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase createtransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase createtransaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createtransaction.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.Write([]byte("Error1 reading body" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	outputDTO, err := h.CreateTransactionUseCase.Execute(&dto)
	if err != nil {
		w.Write([]byte("Error2 reading body" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputDTO)
	if err != nil {
		w.Write([]byte("Error3 reading body" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.Write([]byte("Error4 reading body" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
