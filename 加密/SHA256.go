package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	timestamp := "1641305051"
	secret := "6E18A34A2866CE04651D75708B33727F"

	message := timestamp + "GET" + "/users/self/verify"

	sign := HmacSha256(message, secret)
	fmt.Println(sign)
}

func HmacSha256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
