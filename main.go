package main

import (
	"io"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 8) // eight byte string
	var line string
	line = ""
	for (true) {
		n, err := file.Read(data)
		if (err != nil) {
			if (err == io.EOF) {
				break;
			} else {
				log.Fatal(err)
			}
		}
		slices := strings.Split(string(data[:n]), "\n")
		line += slices[0]
		if len(slices) > 1 {
			fmt.Printf("read: %s\n", line)
			line = ""
			line += slices[len(slices)-1]
		}
		for i := 1; i < len(slices) - 1; i++ {
			fmt.Printf("read: %s\n", slices[i])
		}

	}
	if line != "" {
		fmt.Printf("read: %s\n", line)
	}
	file.Close()
}
