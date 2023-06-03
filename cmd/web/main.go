package main

import (
	"fmt"
	"net/http"

	"github.com/whoajitpatil/reflecton/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "hello world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	// fmt.Printf("Number of bytes written: %d", n)
	// })

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting a web server on %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}