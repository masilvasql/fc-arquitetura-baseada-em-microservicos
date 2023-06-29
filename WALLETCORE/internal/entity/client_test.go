package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	client, err := NewClient("Test Client", "email@email.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, "Test Client", client.Name)
	assert.Equal(t, "email@email.com", client.Email)
	assert.NotEmpty(t, client.CreatedAt)
	assert.NotEmpty(t, client.UpdatedAt)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Test Client", "test@teste.com")
	err := client.Update("Test Client Updated", "testeAtualizado@teste.com")
	assert.Nil(t, err)
	assert.Equal(t, "Test Client Updated", client.Name)
	assert.Equal(t, "testeAtualizado@teste.com", client.Email)
	assert.NotEmpty(t, client.UpdatedAt)
}

func TestUpdateClientWhenARgsAreInvalid(t *testing.T) {
	client, _ := NewClient("Test Client", "test@teste.com")
	err := client.Update("", "")
	assert.NotNil(t, err, "Name is required")
}

func TestAddAccountIntoClient(t *testing.T) {
	client, _ := NewClient("Test Client", "teste@teste.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountIntoClientWhenAccountDoesNotBelongToClient(t *testing.T) {
	client, _ := NewClient("Test Client", "teste@test.com")
	account := NewAccount(client)
	client2, _ := NewClient("Test Client 2", "teste@teste2.com")
	err := client2.AddAccount(account)
	assert.NotNil(t, err)
}
