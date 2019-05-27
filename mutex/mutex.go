package mutex

import (
	"io"
	"net/http"
	"sync"

	"github.com/Ezian/gof-art/utils"
)

// Last image generated
var lastGenerated string
var mutex sync.RWMutex

// HandleGet Retrieve the last generated image
func HandleGet(rw http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	io.WriteString(rw, lastGenerated)
	mutex.RUnlock()
}

// HandlePost Download & Convert & store the image
func HandlePost(rw http.ResponseWriter, req *http.Request) {
	value, err := utils.DownloadAndConvert(rw, req)
	if err == nil {
		mutex.Lock()
		lastGenerated = value
		mutex.Unlock()
	}
}
