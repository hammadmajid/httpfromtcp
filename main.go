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

	var line string

	for {
		var current string

		var bytes [8]byte
		_, err := file.Read(bytes[:])
		if err == io.EOF {
			break
		}

		current = string(bytes[:])
		parts := strings.Split(current, "\n")

		if len(parts) != 1 { // found \n
			line += parts[0]
			fmt.Printf("read: %s\n", line)

			line = parts[1]
		} else {
			line += current
		}
	}

}
