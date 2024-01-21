package core

import (
	"fmt"
	"golang_blockchain_app/crypto"
)

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(privetKey crypto.PrivateKey) error {
	sig, err := privetKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privetKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
