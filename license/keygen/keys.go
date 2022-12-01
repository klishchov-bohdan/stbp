package keygen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/big"
)

type PrivateKey ecdsa.PrivateKey

type pkContainer struct {
	Pub []byte
	D   *big.Int
}

var Curve = elliptic.P384

func NewPrivateKey() (*PrivateKey, error) {
	tmp, err := ecdsa.GenerateKey(Curve(), rand.Reader)
	if err != nil {
		return nil, err
	}
	key := PrivateKey(*tmp)
	return &key, nil
}

func (k *PrivateKey) toEcdsa() *ecdsa.PrivateKey {
	r := ecdsa.PrivateKey(*k)
	return &r
}

func (k PrivateKey) ToBytes() ([]byte, error) {
	ek := k.toEcdsa()
	c := &pkContainer{
		elliptic.Marshal(
			ek.PublicKey.Curve,
			ek.PublicKey.X,
			ek.PublicKey.Y),
		ek.D,
	}
	return toBytes(c)
}

func (k PrivateKey) ToB64String() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (k PrivateKey) ToB32String() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(b), nil
}

func (k PrivateKey) ToHexString() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func PrivateKeyFromBytes(b []byte) (*PrivateKey, error) {
	c := &pkContainer{}
	if err := fromBytes(c, b); err != nil {
		return nil, err
	}
	pk, err := PublicKeyFromBytes(c.Pub)
	if err != nil {
		return nil, err
	}
	k := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey(*pk),
		D:         c.D,
	}
	res := PrivateKey(k)
	return &res, nil
}

func PrivateKeyFromB64String(str string) (*PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}

func PrivateKeyFromB32String(str string) (*PrivateKey, error) {
	b, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}

func PrivateKeyFromHexString(str string) (*PrivateKey, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}


func (k PrivateKey) GetPublicKey() *PublicKey {
	pk := PublicKey(k.PublicKey)
	return &pk
}

type PublicKey ecdsa.PublicKey

func (k *PublicKey) toEcdsa() *ecdsa.PublicKey {
	r := ecdsa.PublicKey(*k)
	return &r
}

func (k PublicKey) ToBytes() []byte {
	
	ek := k.toEcdsa()
	return elliptic.Marshal(ek.Curve, ek.X, ek.Y)
}


func (k PublicKey) ToB64String() string {
	return base64.StdEncoding.EncodeToString(
		k.ToBytes(),
	)
}

func (k PublicKey) ToB32String() string {
	return base32.StdEncoding.EncodeToString(
		k.ToBytes(),
	)
}

func (k PublicKey) ToHexString() string {
	return hex.EncodeToString(
		k.ToBytes(),
	)
}

func PublicKeyFromBytes(b []byte) (*PublicKey, error) {
	x, y := elliptic.Unmarshal(Curve(), b)
	if x == nil {
		return nil, errors.New("invalid key")
	}

	k := ecdsa.PublicKey{
		Curve: Curve(),
		X:     x,
		Y:     y,
	}
	r := PublicKey(k)
	return &r, nil
}

func PublicKeyFromB64String(str string) (*PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}

func PublicKeyFromB32String(str string) (*PublicKey, error) {
	b, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}

func PublicKeyFromHexString(str string) (*PublicKey, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}