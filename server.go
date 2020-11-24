package sia

import "net"

type Handler func(*Reader, *Writer)

func Listen(network, address string, handle Handler) {
	ln, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		reader := NewReader(conn)
		writer := NewWriter(conn)
		go handle(reader, writer)
	}
}
