package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/google/uuid"
)

func encrypt(plaintext *[]byte) (string, error) {
	k := uuid.New()
	keyStr := k.String()
	block, err := aes.NewCipher([]byte(keyStr)[:32])
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, *plaintext, nil)

	*plaintext = ciphertext

	return keyStr, nil
}

func decrypt(encrypted *[]byte, key string) error {
	k := []byte(key)[:32]
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	encVal := *encrypted

	size := gcm.NonceSize()
	nonce := encVal[:size]
	encVal = encVal[size:]

	plaintext, err := gcm.Open(nil, nonce, encVal, nil)

	*encrypted = plaintext

	return err
}
