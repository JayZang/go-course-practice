package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(getCode("Jay")) // 和下一個一樣
	fmt.Println(getCode("Jay"))
	fmt.Println(getCode("Jay1")) // 和上方不一樣
}

// 加鹽 hash
func getCode(message string) string {
	h := hmac.New(sha256.New, []byte("mySecretKey"))
	h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}
