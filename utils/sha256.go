package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

/*
HmacWithShaToBase64

	key char(32)
*/
func HmacWithShaToBase64(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}
