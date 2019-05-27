package utils

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
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

// DownloadAndConvert download image, convert it to ascii, and write the response.
// Return the converted image as string or the error
func DownloadAndConvert(rw http.ResponseWriter, req *http.Request) (result string, err error) {
	decoder := json.NewDecoder(req.Body)
	var data asciiRequest
	err = decoder.Decode(&data)
	if err != nil {
		http.Error(rw, "Wrong json request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Converting image from %q... \n", data.URL)

	start := time.Now()

	image, err := downloadImage(data.URL)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Cannot load or decode image at %q. %v", data.URL, err), http.StatusInternalServerError)
		return
	}

	img, err := ascii.Convert2Ascii(ascii.ScaleImage(*image, data.Width))
	if err != nil {
		http.Error(rw, fmt.Sprintf("Generating Ascii Art failed. %v", err), http.StatusInternalServerError)
		return
	}
	result = fmt.Sprintf(`
	Image from %q
	Generated in %v

	%s
	`, data.URL, time.Since(start), string(img))

	io.WriteString(rw, result)
	return
}

// CreateHandler create a post/get handler with both function
func CreateHandler(posthandler, getHandler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			posthandler(rw, req)
		case http.MethodGet:
			getHandler(rw, req)
		default:
			http.Error(rw, "POST or GET Method are the only allowed", http.StatusMethodNotAllowed)
		}
	}
}
