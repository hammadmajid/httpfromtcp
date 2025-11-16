package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	linesCh := make(chan string)

	go func() {
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

		close(linesCh)
	}()

	return linesCh
}
