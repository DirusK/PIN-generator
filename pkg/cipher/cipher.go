package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func TripleDesEncrypt(data string, cipherKey string) (string, error) {
	key, err := hex.DecodeString(cipherKey)
	if err != nil {
		return "", err
	}
	if len(key) != 24 {
		return "", errors.New("key length must be 24 bytes")
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(data)
	if len(plaintext)%aes.BlockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}

	iv := make([]byte, des.BlockSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encrypted := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, plaintext)

	return hex.EncodeToString(encrypted), nil
}
