package core

import (
	"crypto/sha256"
	"golang_blockchain_app/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(h *Header) types.Hash {
	hash := sha256.Sum256(h.Bytes())
	return types.Hash(hash)
}

type TxHasher struct {
}

func (TxHasher) Hash(tx *Transaction) types.Hash {
	hash := sha256.Sum256(tx.Data)
	return types.Hash(hash)
}
