package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type Message struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature,omitempty"`
}

func EncryptMessage(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plainText, nil), nil
}

func DecryptMessage(key []byte, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, io.ErrUnexpectedEOF
	}
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}

func PerformKeyExchange() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	log.Printf("Generated key size: %d bytes", len(key))
	return key, nil
}
