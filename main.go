package main

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"time"

	"github.com/Ezian/gof-art/ascii"
)

type asciiRequest struct {
	URL   string `json:"url"`
	Width int    `json:"width"`
}

func downloadImage(url string) (*image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func handleASCII(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data asciiRequest
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(rw, "Wrong json request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Converting image from %q... \n", req.URL.Path, data.URL)

	start := time.Now()

	image, err := downloadImage(data.URL)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Cannot load or decode image at %q. %v", data.URL, err), http.StatusInternalServerError)
		return
	}

	result, err := ascii.Convert2Ascii(ascii.ScaleImage(*image, data.Width))
	if err != nil {
		http.Error(rw, fmt.Sprintf("Generating Ascii Art failed. %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, `
	Image from %q
	Generated in %v

	%s
	`, data.URL, time.Since(start), string(result))

}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, you've requested: %s and that's wrong\n", req.URL.Path)
	})
	http.HandleFunc("/ascii", handleASCII)

	http.ListenAndServe(":9999", nil)
}
