package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCounter)  // E.q. /count
	http.HandleFunc("/count/", handlerCounter) // E.q. /count/xxxx
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	// rw.Write([]byte(`My website`))
	fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
}

func handlerCounter(rw http.ResponseWriter, req *http.Request) {
	counter++
	fmt.Fprintf(rw, "Counter = %d\n", counter)
}
