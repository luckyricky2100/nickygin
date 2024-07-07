package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func AESEncrypt(plaintext, key string) (string, error) {
	keyBytes := []byte(key)
	plaintextBytes := []byte(plaintext)

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	plaintextBytes = PKCS5Padding(plaintextBytes, aes.BlockSize)
	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// Combine IV and ciphertext
	encryptedData := append(iv, ciphertext...)

	// Encode the result in base64 for safe storage or transmission
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func AESDecrypt(ciphertext, key string) (string, error) {
	keyBytes := []byte(key)

	// Decode base64-encoded ciphertext
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Extract IV and ciphertext
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	// Create an AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Decrypt the ciphertext
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertextBytes))
	mode.CryptBlocks(plaintext, ciphertextBytes)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext), nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
