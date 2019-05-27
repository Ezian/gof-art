package naive

import (
	"io"
	"net/http"

	"github.com/Ezian/gof-art/utils"
)

// Last image generated
var lastGenerated string

// HandleGet Retrieve the last generated image
func HandleGet(rw http.ResponseWriter, req *http.Request) {
	// Naïve and wrong implementation
	io.WriteString(rw, lastGenerated)
}

// HandlePost Download & Convert & store the image
func HandlePost(rw http.ResponseWriter, req *http.Request) {
	value, err := utils.DownloadAndConvert(rw, req)
	if err == nil {
		lastGenerated = value
	}
}
