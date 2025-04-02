package main

import (
	"fmt"
	"io"
	"net/http"
)

// for both functions we accept two arguments http.ResponseWriter and a *http.Request value.
//
// the *http.Request is used to get information about the request being sent, such as the body being sent in the POST request or info about client

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	io.WriteString(w, "Velcome to my website")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request")
	io.WriteString(w, "hello gang")
}

func main() {

	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", sayHello)

	// err := http.ListenAndServe(":8080", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed \n")
	// } else if err != nil {
	// 	fmt.Printf("Error starting server with port :8080", err)
	// 	os.Exit(1)
	// }

	// introducing a MULTIPLEXER

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}
