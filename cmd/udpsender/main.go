package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "localhost:42069")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		panic(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			panic(err)
		}
	}
}
