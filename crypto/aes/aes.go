package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
	"io"
)

func Encrypt(content string, priKey []byte) (string, *errors.Error) {
	plaintext := []byte(content)
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return strs.EMPTY, errors.Sys("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return strs.EMPTY, errors.Sys("init iv failed: " + err.Error())
	}
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)
	copy(ciphertext[:aes.BlockSize], iv)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(content string, priKey []byte) (string, *errors.Error) {
	ciphertext, err := hex.DecodeString(content)
	if err != nil {
		return strs.EMPTY, errors.Sys("hex.DecodeString(content):" + err.Error())
	}
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return strs.EMPTY, errors.Sys("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	iv := ciphertext[:aes.BlockSize]
	decrypted := make([]byte, len(ciphertext)-aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext[aes.BlockSize:])
	padding := int(decrypted[len(decrypted)-1])
	decrypted = decrypted[:len(decrypted)-padding]
	return string(decrypted), nil
}
