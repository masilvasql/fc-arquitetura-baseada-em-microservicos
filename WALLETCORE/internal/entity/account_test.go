package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {

	client, err := NewClient("Test Client", "teste@teste.com")

	assert.Nil(t, err)

	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.NotEmpty(t, account.ID)

	assert.Equal(t, 0.0, account.Balance)
	assert.NotEmpty(t, account.CreatedAt)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestCreateNewAccountWhenClientIsNil(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditIntoAccount(t *testing.T) {
	client, err := NewClient("Test Client", "teste@teste.com")
	assert.Nil(t, err)

	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)
	assert.NotEmpty(t, account.UpdatedAt)

}

func TestDebitIntoAccount(t *testing.T) {
	client, err := NewClient("Test Client", "teste@teste.com")
	assert.Nil(t, err)

	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)

	account.Debit(50.0)
	assert.Equal(t, 50.0, account.Balance)

	assert.NotEmpty(t, account.UpdatedAt)

}
