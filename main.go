package main

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am message1")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/message1", messageHandler)

	http.ListenAndServe(":8080", mux)

}
