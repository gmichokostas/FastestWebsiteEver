package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var (
	port = "80"
	hdrs = "HTTP/1.1 200 k\nContent-Length: %d\ncontent-encoding: deflate\n\n"
)

func checkError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return true
	}
	return false
}

func handleClient(conn net.Conn) {
	var content []byte

	file, err := ioutil.ReadFile("index.html")
	checkError(err)

	headers := fmt.Sprintf(hdrs, len(file))
	content = append([]byte(headers), file...)

	_, err = conn.Write(content)
	checkError(err)

	conn.Close()
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := ln.Accept()
		if checkError(err) {
			break
		}

		go handleClient(conn)
	}
}
