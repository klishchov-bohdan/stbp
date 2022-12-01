package keygen

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type License struct {
	Data []byte
	R    *big.Int
	S    *big.Int
}

func NewLicense(k *PrivateKey, data []byte) (*License, error) {
	l := &License{
		Data: data,
	}

	if h, err := l.hash(); err != nil {
		return nil, err
	} else if r, s, err := ecdsa.Sign(rand.Reader, k.toEcdsa(), h); err != nil {
		return nil, err
	} else {
		l.R = r
		l.S = s
	}
	return l, nil
}

func (l *License) hash() ([]byte, error) {
	h256 := sha256.New()

	if _, err := h256.Write(l.Data); err != nil {
		return nil, err
	}
	return h256.Sum(nil), nil
}

func (l *License) Verify(k *PublicKey) (bool, error) {
	h, err := l.hash()
	if err != nil {
		return false, err
	}
	return ecdsa.Verify(k.toEcdsa(), h, l.R, l.S), nil
}

func (l *License) ToBytes() ([]byte, error) {
	return toBytes(l)
}

func (l *License) ToB64String() (string, error) {
	return toB64String(l)
}

func (l *License) ToB32String() (string, error) {
	return toB32String(l)
}

func (l *License) ToHexString() (string, error) {
	return toHexString(l)
}

func LicenseFromBytes(b []byte) (*License, error) {
	l := &License{}
	return l, fromBytes(l, b)
}

func LicenseFromB64String(str string) (*License, error) {
	l := &License{}
	return l, fromB64String(l, str)
}

func LicenseFromB32String(str string) (*License, error) {
	l := &License{}
	return l, fromB32String(l, str)
}

func LicenseFromHexString(str string) (*License, error) {
	l := &License{}
	return l, fromHexString(l, str)
}