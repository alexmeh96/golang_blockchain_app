package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeypairSignVerifySuccess(t *testing.T) {
	privetKey := GeneratePrivateKey()
	publicKey := privetKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privetKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(publicKey, msg))
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privetKey := GeneratePrivateKey()
	publicKey := privetKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privetKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivetKey := GeneratePrivateKey()
	otherPublicKey := otherPrivetKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("test msg")))
}
