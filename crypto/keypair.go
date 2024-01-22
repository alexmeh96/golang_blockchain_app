package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"golang_blockchain_app/types"
	"math/big"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{r, s}, nil
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		Key: &k.key.PublicKey,
	}
}

type PublicKey struct {
	Key *ecdsa.PublicKey
}

// GobEncode implementation of the GomEncoder interface,
// for correct encoding of the Key *ecdsa.PublicKey field
func (k *PublicKey) GobEncode() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(k.Key.X); err != nil {
		return nil, err
	}
	if err := encoder.Encode(k.Key.Y); err != nil {
		return nil, err
	}
	if err := encoder.Encode(k.Key.Curve.Params().Name); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// GobDecode implementation of the GobDecoder interface,
// for correct encoding of the Key *ecdsa.PublicKey field
func (k *PublicKey) GobDecode(data []byte) error {
	// Decode the necessary fields of *ecdsa.PublicKey
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	var x, y big.Int
	var curveName string

	if err := decoder.Decode(&x); err != nil {
		return err
	}
	if err := decoder.Decode(&y); err != nil {
		return err
	}
	if err := decoder.Decode(&curveName); err != nil {
		return err
	}

	curve := elliptic.P256()
	if curveName != curve.Params().Name {
		return fmt.Errorf("invalid curve name: %s", curveName)
	}

	k.Key = &ecdsa.PublicKey{
		X:     &x,
		Y:     &y,
		Curve: curve,
	}

	return nil
}

func (k *PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key, k.Key.X, k.Key.Y)
}

func (k *PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	R, S *big.Int
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.Key, data, sig.R, sig.S)
}
