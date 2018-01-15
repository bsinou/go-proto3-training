// Simple examples to test the net/http library

package nethttp

import (
	//	"fmt"
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm + "\n"))
}

func launchserver() {
	//	fmt.Println("Registered Hello World handler on port 4000")

	// Start a server on port 3000 in another thread
	go func() {
		mux := http.NewServeMux()
		th := &timeHandler{format: time.RFC1123}
		mux.Handle("/time", th)

		h := &helloer{}
		mux.Handle("/hello", h)

		log.Println("Listening on port 3000 ...")

		http.ListenAndServe(":3000", mux)
	}()

	// start a second one on port 4000
	// register the HelloWorld hndle function on the default ServeMux
	http.HandleFunc("/foo", HelloWorld)
	log.Println("Listening on port 4000 ...")
	// Listen with the default Mux
	http.ListenAndServe(":4000", nil)

}

type helloer struct{}

func (h *helloer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!\n"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!\n"))
}
