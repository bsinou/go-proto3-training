// Canonical example to easily start a server that replies "hello world" on port 7070
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":7070", nil)
}
