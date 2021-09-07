package main

import (
	"net/http"

	"github.com/jutinapas/webservice/controllers"
)

func main() {
	// fmt.Println("Hello, World!")

	// port := 3000
	// port, err := startWebServer(port, 2)
	// fmt.Println(port, err)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

// // data type of arguments can be defined as a comma-delimited
// // multiple return values is allowed in Go
// func startWebServer(port, numberOfRetries int) (int, error) {
// 	fmt.Println("starting server...")
// 	fmt.Println("starting started on port", port)
// 	fmt.Println("number of retries", numberOfRetries)

// 	// error is a pointer data type
// 	return port, nil
// }
