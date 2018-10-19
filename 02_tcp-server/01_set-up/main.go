package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 建立 tcp server 並監聽 8080 port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		// 等待用戶連線
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	// 取得客戶端傳輸資料
	ns := bufio.NewScanner(conn)
	for ns.Scan() {
		fmt.Println(ns.Text())
	}

	// 當連線中斷時才會執行到這
	fmt.Println("Code got here.")
}
