package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func DecodeMD5(value string) string {
	result, err := hex.DecodeString(value)
	if err != nil {
		return ""
	}

	return string(result)
}
