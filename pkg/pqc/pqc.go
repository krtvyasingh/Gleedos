package pqc

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type KyberKeypair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func GenerateKyberKeypair() (*KyberKeypair, error) {
	pub := make([]byte, 1568)
	priv := make([]byte, 3168)
	if _, err := io.ReadFull(rand.Reader, pub); err != nil {
		return nil, err
	}
	if _, err := io.ReadFull(rand.Reader, priv); err != nil {
		return nil, err
	}
	return &KyberKeypair{PublicKey: pub, PrivateKey: priv}, nil
}

func Encapsulate(publicKey []byte) ([]byte, []byte, error) {
	ciphertext := make([]byte, 1568)
	sharedSecret := sha256.Sum256(publicKey[:32])
	return ciphertext, sharedSecret[:], nil
}
