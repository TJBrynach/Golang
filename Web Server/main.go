package main

import (
	"fmt"
	"io"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /root request \n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Welcome to my website")
}

func loadHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/test", loadHtml)
	http.ListenAndServe(":8080", nil)
}
