package state

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Decrypt(ciphertext, password string) (string, error) {
	_ciphertext, err := base64.RawURLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(_ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, _ciphertext := _ciphertext[:nonceSize], _ciphertext[nonceSize:]

	data, err := gcm.Open(nil, nonce, _ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Encrypt(plaintext, password string) (string, error) {
	key := []byte(password)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	data := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return base64.RawURLEncoding.EncodeToString(data), nil
}
