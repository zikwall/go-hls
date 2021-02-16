package io

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Listening to the input stream of the process
func FromStdin() io.Reader {
	return os.Stdin
}

// Listening to the connection over the TCP protocol
// You can simulate TCP data via a terminal:
//
// 	```bash
// 		$ nc localhost 1339
// 	```
func FromTCP(port ...int) io.Reader {
	p := 1339

	if len(port) > 0 && port[0] > 0 {
		p = port[0]
	}

	ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", p))
	conn, _ := ln.Accept()

	return conn
}
