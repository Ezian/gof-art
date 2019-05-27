package channel

import (
	"io"
	"net/http"

	"github.com/Ezian/gof-art/utils"
)

// Last image generated
var (
	postedImage = make(chan string)
	gettedImage = make(chan string)
)

// Cache manager
func init() {
	go func() {
		var lastGenerated string
		for {
			select {
			// New posted image
			case lastGenerated = <-postedImage:
				// No new image : we update the read channel
			case gettedImage <- lastGenerated:
			}
		}

	}()
}

// HandleGet Retrieve the last generated image
func HandleGet(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, <-gettedImage)
}

// HandlePost Download & Convert & store the image
func HandlePost(rw http.ResponseWriter, req *http.Request) {
	value, err := utils.DownloadAndConvert(rw, req)
	if err == nil {
		postedImage <- value
	}
}
