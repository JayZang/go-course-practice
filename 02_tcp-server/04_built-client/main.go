package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 連線至 server 端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Hello, I am Jay")

	// 持續監聽來自 server 的回覆
	io.Copy(os.Stdout, conn)
}
