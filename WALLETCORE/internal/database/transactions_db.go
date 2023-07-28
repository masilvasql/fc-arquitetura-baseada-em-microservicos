package database

import (
	"context"
	"database/sql"
	"github.com/masilvasql/fc-ms-wallet/pkg/uow"

	"github.com/masilvasql/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func (t *TransactionDB) Register(name string, fc uow.RepositoryFactory) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionDB) GetRepository(ctx context.Context, name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionDB) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionDB) CommitOrRollback() error {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionDB) Rollback() error {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionDB) UnRegister(name string) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{DB: db}
}

func (t *TransactionDB) Create(transcation *entity.Transaction) error {
	stmt, err := t.DB.Prepare(`
		INSERT INTO transactions (
			id,
			account_id_from,
			account_id_to,
			amount,
			created_at
		) VALUES (?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		transcation.ID,
		transcation.AccountFrom.ID,
		transcation.AccountTo.ID,
		transcation.Amount,
		transcation.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
