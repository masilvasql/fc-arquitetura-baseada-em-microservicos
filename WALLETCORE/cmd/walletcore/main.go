package main

import (
	"context"
	"database/sql"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masilvasql/fc-ms-wallet/internal/database"
	"github.com/masilvasql/fc-ms-wallet/internal/event"
	"github.com/masilvasql/fc-ms-wallet/internal/event/handler"
	createaccount "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_account"
	createclient "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/masilvasql/fc-ms-wallet/internal/web"
	"github.com/masilvasql/fc-ms-wallet/internal/web/webserver"
	"github.com/masilvasql/fc-ms-wallet/pkg/events"
	"github.com/masilvasql/fc-ms-wallet/pkg/kafka"
	"github.com/masilvasql/fc-ms-wallet/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTrasactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	unitOfWork := uow.NewUow(ctx, db)
	unitOfWork.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	unitOfWork.Register("ClientDB", func(tx *sql.Tx) interface{} {
		return database.NewClientDB(db)
	})

	unitOfWork.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	creaTeTransactionUseCase := createtransaction.NewCreateTransactionUseCase(unitOfWork, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webServer := webserver.NewWebServer(":8080")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*creaTeTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	err = webServer.Start()
	if err != nil {
		panic(err)
	}

}
