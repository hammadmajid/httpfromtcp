package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		println("Failed to open file")
		os.Exit(1)
	}

	for {
		var bytes [8]byte
		_, err := file.Read(bytes[:])
		if err != nil {
			os.Exit(0)
		}

		fmt.Printf("read: %s\n", bytes)
	}
}
