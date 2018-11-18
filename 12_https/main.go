package main

import (
	"fmt"
	"log"
	"net/http"
)

// 生成未認證憑證及私鑰
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=yourdomainname
// 認證之憑證可於 LetEncrypt 申請
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Success!")
}
