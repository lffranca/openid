package entities

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"math/big"
)

type Key struct {
	Kty    string   `json:"kty"`
	Use    string   `json:"use"`
	Kid    string   `json:"kid"`
	X5T    string   `json:"x5t"`
	N      string   `json:"n"`
	E      string   `json:"e"`
	X5C    []string `json:"x5c"`
	Issuer string   `json:"issuer"`
}

func (key *Key) ToRSA() (*rsa.PublicKey, error) {
	decodedE, err := base64.RawURLEncoding.DecodeString(key.E)
	if err != nil {
		return nil, err
	}

	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}

	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}

	decodedN, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		return nil, err
	}

	pubKey.N.SetBytes(decodedN)
	return pubKey, nil
}

func (key *Key) ToBytes() ([]byte, error) {
	rsaKey, errRsa := key.ToRSA()
	if errRsa != nil {
		return nil, errRsa
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(rsaKey)
	if err != nil {
		return nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}
