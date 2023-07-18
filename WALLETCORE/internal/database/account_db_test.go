package database

import (
	"database/sql"
	"testing"

	"github.com/masilvasql/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec(`CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance int, created_at date);`)
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Marcelo", "teste@teste.com")
}

func (s *AccountDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestFindById() {

	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)

	accountFound, err := s.accountDB.FindById(account.ID)
	s.Nil(err)
	s.Equal(accountFound.ID, account.ID)
	s.Equal(accountFound.Client.ID, account.Client.ID)
	s.Equal(accountFound.Client.Name, account.Client.Name)
	s.Equal(accountFound.Client.Email, account.Client.Email)
	s.Equal(accountFound.Balance, account.Balance)
	s.Equal(s.client.ID, accountFound.Client.ID)

}
