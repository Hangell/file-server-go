package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <directory> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1] // position 0 is the file name
	port := os.Args[2]

	authenticator := NewAuthenticator()
	StartServer(httpDir, port, authenticator)
}
