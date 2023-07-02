package database

import (
	"database/sql"
	"testing"

	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec(`CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance int, created_at date);`)
	db.Exec(`CREATE TABLE transactions(id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount real, created_at date);`)
	s.transactionDB = NewTransactionDB(db)
	s.client, err = entity.NewClient("Marcelo", "teste@teste.com")
	s.Nil(err)
	s.client2, _ = entity.NewClient("client 2", "client2@teste.com")
	s.Nil(err)
	s.accountFrom = entity.NewAccount(s.client)
	s.accountFrom.Balance = 1000
	s.accountTo = entity.NewAccount(s.client2)
}

func (s *TransactionDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

// IRA RODAR TODOS OS TESTS SUITES
func TestTransactionDbTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100.0)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)

}
