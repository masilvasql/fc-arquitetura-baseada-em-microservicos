package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, err := NewClient("Test Client 1", "teste@teste.com")
	assert.Nil(t, err)

	accountFrom := NewAccount(client1)
	accountFrom.Credit(100.0)

	client2, err := NewClient("Test Client 2", "teste2@teste.com")
	assert.Nil(t, err)
	accountTo := NewAccount(client2)

	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.NotEmpty(t, transaction.ID)
	assert.NotEmpty(t, transaction.CreatedAt)
	assert.Equal(t, 0.0, accountFrom.Balance)
	assert.Equal(t, 100.0, accountTo.Balance)
}

func TestCreateTransactionWithInsuficientFunds(t *testing.T) {
	client1, err := NewClient("Test Client 1", "t1@teste.com")
	assert.Nil(t, err)

	accountFrom := NewAccount(client1)
	accountFrom.Credit(100.0)

	client2, err := NewClient("Test Client 2", "teste2@teste.com")
	assert.Nil(t, err)
	accountTo := NewAccount(client2)

	transaction, err := NewTransaction(accountFrom, accountTo, 200)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "Insufficient funds", err.Error())
	assert.Equal(t, 100.0, accountFrom.Balance)
	assert.Equal(t, 0.0, accountTo.Balance)
}

func TestCreateTransactionWithInvalidAmount(t *testing.T) {
	client1, err := NewClient("Test Client 1", "t1@teste.com")
	assert.Nil(t, err)

	accountFrom := NewAccount(client1)
	accountFrom.Credit(100.0)

	client2, err := NewClient("Test Client 2", "t2@teste.com")
	assert.Nil(t, err)
	accountTo := NewAccount(client2)

	transaction, err := NewTransaction(accountFrom, accountTo, -100)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "Amount is required", err.Error())
	assert.Equal(t, 100.0, accountFrom.Balance)
	assert.Equal(t, 0.0, accountTo.Balance)
}
