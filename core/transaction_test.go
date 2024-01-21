package core

import (
	"github.com/stretchr/testify/assert"
	"golang_blockchain_app/crypto"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privateKey))
	assert.Nil(t, tx.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivateKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}
