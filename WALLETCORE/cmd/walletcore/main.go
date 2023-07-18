package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masilvasql/fc-ms-wallet/internal/database"
	"github.com/masilvasql/fc-ms-wallet/internal/event"
	createaccount "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_account"
	createclient "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/masilvasql/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/masilvasql/fc-ms-wallet/internal/web"
	"github.com/masilvasql/fc-ms-wallet/internal/web/webserver"
	"github.com/masilvasql/fc-ms-wallet/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTrasactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	creaTeTransactionUseCase := createtransaction.CreateNewTranasction(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(":3000")
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
