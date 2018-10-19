package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	// http 請求回覆後即斷開連接
	defer conn.Close()

	request(conn)
	response(conn)
}

// 解析請求
func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
}

// 回覆請求
func response(conn net.Conn) {
	body := `
		<html>
			<head></head>
			<body>
				<div>HelloYa</div>
			</body>
		</html>
	`

	// 送出 http 協議之指令
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
