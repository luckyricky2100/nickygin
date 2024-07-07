package main

import (
	"testing"

	"nickygin.com/pkg/util"
)

func TestAEC(t *testing.T) {
	plaintext := "Hello, AES!Hello, AES!Hello, AES!Hello, AES!Hello, AES!"
	key := "my-secret-key123"

	// Encrypt
	encryptedData, err := util.AESEncrypt(plaintext, key)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// Decrypt
	decryptedText, err := util.AESDecrypt(encryptedData, key)
	if err != nil {
		t.Errorf(err.Error())
	}
	if decryptedText != plaintext {
		t.Errorf("%s does not equal %s", decryptedText, plaintext)
	}
}
