package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var b []byte = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%x", b)
	}
}
