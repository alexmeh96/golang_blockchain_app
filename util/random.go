package util

import (
	"golang_blockchain_app/core"
	"golang_blockchain_app/crypto"
	"golang_blockchain_app/types"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() types.Hash {
	return types.HashFromBytes(RandomBytes(32))
}

// NewRandomTransaction return a new random transaction whithout signature.
func NewRandomTransaction(size int) *core.Transaction {
	return core.NewTransaction(RandomBytes(size))
}

func NewRandomTransactionWithSignature(t *testing.T, privKey crypto.PrivateKey, size int) *core.Transaction {
	tx := NewRandomTransaction(size)
	assert.Nil(t, tx.Sign(privKey))
	return tx
}
