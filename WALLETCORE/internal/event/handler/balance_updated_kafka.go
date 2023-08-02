package handler

import (
	"fmt"
	"github.com/masilvasql/fc-ms-wallet/pkg/events"
	"github.com/masilvasql/fc-ms-wallet/pkg/kafka"
	"sync"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{Kafka: kafka}
}

func (h *BalanceUpdatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("BalanceUpdatedKafkaHandler called", message.GetPayload())
}
