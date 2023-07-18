package web

import (
	"encoding/json"
	createaccount "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_account"
	"net/http"
)

type WebAccountHandler struct {
	CreateAccountUseCase createaccount.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase createaccount.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUseCase: createAccountUseCase,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto createaccount.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	outputDTO, err := h.CreateAccountUseCase.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
