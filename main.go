package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	port := 3000

	port, err := startWebServer(port, 2)
	fmt.Println(port, err)
}

// data type of arguments can be defined as a comma-delimited
// multiple return values is allowed in Go
func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("starting server...")
	fmt.Println("starting started on port", port)
	fmt.Println("number of retries", numberOfRetries)

	// error is a pointer data type
	return port, nil
}
