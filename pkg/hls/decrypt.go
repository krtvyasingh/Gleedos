package hls

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAES128(data, key, iv []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("invalid AES-128 key length")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != 16 {
		iv = make([]byte, 16)
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	mode.CryptBlocks(decrypted, data)
	return decrypted, nil
}
