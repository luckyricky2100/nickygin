package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Hmac(key, data string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func HmacSha256(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
