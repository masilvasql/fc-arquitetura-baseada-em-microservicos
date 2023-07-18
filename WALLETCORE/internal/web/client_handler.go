package web

import (
	"encoding/json"
	createclient "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_client"
	"net/http"
)

type WebClientHandler struct {
	CreateClientUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto createclient.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error1 reading body" + err.Error()))
		return
	}
	outputDTO, err := h.CreateClientUseCase.Execute(&dto)
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

	//write json on body
	if err != nil {
		w.Write([]byte("Error4 reading body" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.Write([]byte("Error5 reading body" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
