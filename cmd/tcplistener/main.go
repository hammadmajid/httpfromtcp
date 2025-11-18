package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Listening on " + conn.RemoteAddr().String())

		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Printf("%s\n", line)
		}
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	linesCh := make(chan string)

	go func() {
		defer close(linesCh)

		var buffer string
		for {
			var current string

			var bytes [8]byte
			_, err := f.Read(bytes[:])
			if err == io.EOF {
				break
			}

			current = string(bytes[:])
			parts := strings.Split(current, "\n")

			if len(parts) != 1 { // found \n
				buffer += parts[0]
				linesCh <- buffer

				buffer = parts[1]
			} else {
				buffer += current
			}
		}
	}()

	return linesCh
}
