package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type asciiRequest struct {
	URL string `json:"url"`
}

func handleASCII(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data asciiRequest
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, "Converting image from %q... \n", req.URL.Path, data.URL)

	// TODO : Download image from URL
	// TODO : Convert image to ASCII
	// TODO : Return ascii as a response
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, you've requested: %s and that's wrong\n", req.URL.Path)
	})
	http.HandleFunc("/ascii", handleASCII)

	http.ListenAndServe(":9999", nil)
}
